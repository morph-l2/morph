package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strconv"
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
	"github.com/scroll-tech/go-ethereum/core"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/crypto/bls12381"
	"github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/params"
	"github.com/tendermint/tendermint/blssignatures"
)

const (
	txSlotSize     = 32 * 1024
	txMaxSize      = 4 * txSlotSize // 128KB
	minFinalizeNum = 2              // min finalize num from contract
	minGasLimit    = 1000
)

type SR struct {
	L1Client       iface.Client
	L2Clients      []iface.L2Client
	Rollup         iface.IRollup
	L2Submitter    iface.ISubmitter
	rollupAddr     common.Address
	secondInterval time.Duration
	ctx            context.Context
	chainId        *big.Int
	privKey        *ecdsa.PrivateKey
	metrics        *metrics.Metrics
	abi            *abi.ABI
	batchPool      map[uint64]bindings.IRollupBatchData
	txTimout       time.Duration
	Finalize       bool
	MaxFinalizeNum uint64
	PriorityRollup bool
}

func NewSR(ctx context.Context, l1 iface.Client, l2 []iface.L2Client, rollup iface.IRollup, interval time.Duration, chainId *big.Int, priKey *ecdsa.PrivateKey, rollupAddr common.Address, metrics *metrics.Metrics, abi *abi.ABI, txTimeout time.Duration, maxBlock, minBlock uint64, l2Submitter iface.ISubmitter, finalize bool, maxFinalizeNum uint64, priorityRollup bool) *SR {

	return &SR{
		ctx:            ctx,
		L1Client:       l1,
		L2Clients:      l2,
		Rollup:         rollup,
		secondInterval: interval,
		privKey:        priKey,
		chainId:        chainId,
		rollupAddr:     rollupAddr,
		metrics:        metrics,
		abi:            abi,
		batchPool:      make(map[uint64]bindings.IRollupBatchData),
		txTimout:       txTimeout,
		L2Submitter:    l2Submitter,
		Finalize:       finalize,
		MaxFinalizeNum: maxFinalizeNum,
		PriorityRollup: priorityRollup,
	}
}

func (sr *SR) Start() {
	// block node startup during initial sync and print some helpful logs
	t := time.NewTicker(sr.secondInterval)
	defer t.Stop()

	// metrics
	metrics_query_tick := time.NewTicker(time.Second * 5)
	defer metrics_query_tick.Stop()
	go func() {

		for {
			select {
			case <-sr.ctx.Done():
				return
			case <-metrics_query_tick.C:
				// get last finalized
				lastFinalized, err := sr.Rollup.LastFinalizedBatchIndex(nil)
				if err != nil {
					log.Error("get last finalized error", "error", err)
					continue
				}
				// get last committed
				lastCommited, err := sr.Rollup.LastCommittedBatchIndex(nil)
				if err != nil {
					log.Error("get last committed error", "error", err)
					continue
				}
				sr.metrics.SetLastFinalizedBatchIndex(lastFinalized.Uint64())
				sr.metrics.SetLastCommittedBatchIndex(lastCommited.Uint64())
				sr.metrics.SetLastFinalizedCommitedBatchIndexDiff(lastCommited.Uint64() - lastFinalized.Uint64())

				// get last rolluped l2 block
				l2LatestBlockNumberRolluped, err := sr.Rollup.LatestL2BlockNumber(nil)
				if err != nil {
					log.Error("get last l2 block number error", "error", err)
					continue
				}
				sr.metrics.SetLastL2BlockNumberRolluped(l2LatestBlockNumberRolluped.Uint64())
				// get last l2 block
				// todo: get the largest block number from all l2 clients
				l2BlockNumber, err := sr.L2Clients[0].BlockNumber(context.Background())
				if err != nil {
					log.Error("get l2 block number error", "error", err)
					continue
				}
				sr.metrics.SetL2BlockNumber(l2BlockNumber)

				// diff
				sr.metrics.SetL2BlockNumberDiff(l2BlockNumber - l2LatestBlockNumberRolluped.Uint64())

				// get balacnce of wallet
				balance, err := sr.L1Client.BalanceAt(context.Background(), crypto.PubkeyToAddress(sr.privKey.PublicKey), nil)
				if err != nil {
					log.Error("get wallet balance error", "error", err)
					continue
				}
				// balance to eth
				balanceEth := new(big.Rat).SetFrac(balance, big.NewInt(params.Ether))

				// parse float64 from string
				balanceEthFloat, err := strconv.ParseFloat(balanceEth.FloatString(18), 64)
				if err != nil {
					log.Warn("parse balance to float error", "error", err)
					continue
				}

				sr.metrics.SetWalletBalance(balanceEthFloat)
			}
		}
	}()

	for {
		if err := sr.rollup(); err != nil {
			if utils.IsRpcErr(err) {
				sr.metrics.IncRpcErrors()
			}
			time.Sleep(2 * time.Second)
			log.Error("rollup failed,wait for the next try", "error", err)
		}
		if sr.Finalize {
			if err := sr.finalize(); err != nil {
				log.Error("finalize failed", "error", err)
			}
			time.Sleep(5 * time.Second)
		}
	}
}

func (sr *SR) finalize() error {
	log.Info("check finalize")
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

	// check if need to finalize
	// get last
	period, err := sr.Rollup.FINALIZATIONPERIODSECONDS(nil)
	if err != nil {
		return fmt.Errorf("get finalization period error:%v", err)
	}
	periodU64 := period.Uint64()
	var finalizeIndex uint64

	for i := lastCommited.Uint64(); i > lastFinalized.Uint64(); i-- {
		committedBatch, err := sr.Rollup.CommittedBatchStores(nil, big.NewInt(int64(i)))
		if err != nil {
			return fmt.Errorf("get committed batch error:%v", err)
		}
		// timestamp
		timestamp := uint64(time.Now().Unix())
		if committedBatch.OriginTimestamp.Uint64()+periodU64 < timestamp {
			finalizeIndex = i
			break
		}
	}

	if finalizeIndex == 0 {
		log.Info("no need to finalize")
		return nil
	}

	finalizeCnt := finalizeIndex - lastFinalized.Uint64()
	if finalizeCnt < minFinalizeNum {
		log.Info(fmt.Sprintf("no need to finalize, finalize num < %d", minFinalizeNum))
		return nil
	}
	// finalize up to 5 batch
	if finalizeCnt > sr.MaxFinalizeNum {
		finalizeCnt = sr.MaxFinalizeNum
	}
	// send tx
	// update gas limit
	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}
	opts.NoSend = true
	tx, err := sr.Rollup.FinalizeBatchsByNum(opts, big.NewInt(int64(finalizeCnt)))
	if err != nil {
		return fmt.Errorf("craft FinalizeBatchsByNum tx failed:%v", err)
	}
	if uint64(tx.Size()) > txMaxSize {
		return core.ErrOversizedData
	}
	newTx, err := UpdateGasLimit(tx)
	if err != nil {
		return fmt.Errorf("update gas limit error:%v", err)
	}
	newSignedTx, err := opts.Signer(opts.From, newTx)
	if err != nil {
		return fmt.Errorf("sign tx error:%v", err)
	}

	var receipt *types.Receipt
	err = sr.L1Client.SendTransaction(context.Background(), newSignedTx)
	if err != nil {
		log.Info("send tx error", "error", err.Error())
		// ErrReplaceUnderpriced,ErrAlreadyKnown
		if utils.ErrStringMatch(err, core.ErrReplaceUnderpriced) ||
			utils.ErrStringMatch(err, core.ErrAlreadyKnown) {
			receipt, newSignedTx, err = sr.replaceTx(newSignedTx)
			if err != nil {
				return fmt.Errorf("replace tx error:%v", err)
			} else {
				log.Info("replace tx success")
			}
		} else if utils.ErrStringMatch(err, core.ErrGasLimit) { // ErrGasLimit
			log.Error("tx exceeds block gas limit", "gas", newSignedTx.Gas(), "finalize_cnt", finalizeCnt)
			return fmt.Errorf("send tx error:%v", err.Error())
		} else if utils.ErrStringMatch(err, core.ErrNonceTooLow) { //ErrNonceTooLow
			return fmt.Errorf("send tx error:%v", err.Error())
		} else {
			return fmt.Errorf("send tx error:%v", err.Error())
		}
	} else {
		log.Info("tx sent",
			// for business
			"finalize_cnt", finalizeCnt,
			// tx
			"tx_hash", newSignedTx.Hash().String(),
			"type", newSignedTx.Type(),
			"gas", newSignedTx.Gas(),
			"nonce", newSignedTx.Nonce(),
			"size", newSignedTx.Size(),
			"gas_price", newSignedTx.GasPrice().String(),
			"tip", newSignedTx.GasTipCap().String(),
			"fee_cap", newSignedTx.GasFeeCap().String(),
		)
	}

	// wait receipt
	if receipt == nil {
		receipt, err = sr.waitReceiptWithTimeout(sr.txTimout, newSignedTx.Hash())
		if err != nil {
			log.Error("wait receipt error", "error", err.Error())
			if utils.ErrStringMatch(err, errors.New("timeout")) {
				receipt, newSignedTx, err = sr.replaceTx(newSignedTx)
				if err != nil {
					return fmt.Errorf("replace tx error:%v", err)
				} else {
					log.Info("replace tx success")
				}
			} else {
				return fmt.Errorf("wait receipt error:%v", err)
			}
		}
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("finalize tx failed",
			//block info
			"block_number", receipt.BlockNumber.Uint64(),
			// for business
			"finalize_cnt", finalizeCnt,
			// receipt info
			"gas_used", receipt.GasUsed,
			"tx_hash", newSignedTx.Hash().String(),
			// tx info
			"tx_hash", newSignedTx.Hash().String(),
			"type", newSignedTx.Type(),
			"gas", newSignedTx.Gas(),
			"nonce", newSignedTx.Nonce(),
			"size", newSignedTx.Size(),
			"gas_price", newSignedTx.GasPrice().String(),
			"tip", newSignedTx.GasTipCap().String(),
			"fee_cap", newSignedTx.GasFeeCap().String(),
		)
		return fmt.Errorf("tx failed")
	} else {

		gasFee := newSignedTx.GasPrice().Uint64() * receipt.GasUsed
		gasFeeEther := new(big.Rat).SetFrac(big.NewInt(int64(gasFee)), big.NewInt(params.Ether))
		log.Info("finalize tx success",
			//block info
			"block_number", receipt.BlockNumber.Uint64(),
			// for business
			"finalize_cnt", finalizeCnt,
			// receipt info
			"gas_used", receipt.GasUsed,
			// tx info
			"tx_hash", newSignedTx.Hash().String(),
			"type", newSignedTx.Type(),
			"gas", newSignedTx.Gas(),
			"nonce", newSignedTx.Nonce(),
			"size", newSignedTx.Size(),
			"gas_price", newSignedTx.GasPrice().String(),
			"tip", newSignedTx.GasTipCap().String(),
			"fee_cap", newSignedTx.GasFeeCap().String(),
			"gas_fee", gasFeeEther.FloatString(18),
		)
	}
	return nil

}

func (sr *SR) rollup() error {

	if !sr.PriorityRollup {
		// is the turn of the submitter
		nextSubmitter, _, _, err := sr.L2Submitter.GetNextSubmitter(nil)
		if err != nil {
			return fmt.Errorf("get next submitter error:%v", err)
		}

		if nextSubmitter.Hex() == sr.walletAddr() {
			log.Info("submitter is me, start to submit")
		} else {
			log.Info("submitter is not me, wait for the next turn")
			return nil
		}
	}

	nonce, err := sr.L1Client.NonceAt(context.Background(), crypto.PubkeyToAddress(sr.privKey.PublicKey), nil)
	if err != nil {
		return fmt.Errorf("query layer1 nonce error:%v", err.Error())
	}
	batchIndex, err := sr.Rollup.LastCommittedBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last committed batch index error:%v", err)
	}

	index := batchIndex.Uint64() + 1
	log.Info("batch info", "last_commit_batch", batchIndex.Uint64(), "batch_will_get", index)
	batch, err := GetRollupBatchByIndex(index, sr.L2Clients)
	if err != nil {
		return fmt.Errorf("get rollup batch by index err:%v", err)
	}

	// check if the batch is valid
	if batch == nil {
		time.Sleep(sr.secondInterval)
		log.Info("new batch not found, wait for the next turn")
		return nil
	}

	// log batch
	// batchJson, err := json.Marshal(batch)
	// if err != nil {
	// 	return err
	// }
	// log.Info("batch info", "batch_index", index, "batch_json", string(batchJson))

	if len(batch.Signatures) == 0 {
		log.Info("length of batch signature is empty, wait for the next turn")
		time.Sleep(time.Second * 3)
		return nil
	}

	var chunks [][]byte
	// var blobChunk []byte
	for _, chunk := range batch.Chunks {
		chunks = append(chunks, chunk)
		// blobChunk = append(blobChunk, chunk...)
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

	calldata, err := sr.abi.Pack("commitBatch", rollupBatch, uint32(minGasLimit))
	if err != nil {
		return fmt.Errorf("pack calldata error:%v", err)
	}
	// var signedTx *types.Transaction

	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	opts.NoSend = true
	opts.Nonce = big.NewInt(int64(nonce))

	if err != nil {
		return fmt.Errorf("dial ethclient error:%v", err)
	}

	// versioned hashes

	tip, gasFeeCap, err := sr.GetGasTipAndCap()
	if err != nil {
		return fmt.Errorf("get gas tip and cap error:%v", err)

	}

	var tx *types.Transaction
	// blob tx
	if batch.Sidecar.Blobs == nil || len(batch.Sidecar.Blobs) == 0 {
		tx, err = sr.Rollup.CommitBatch(opts, rollupBatch, uint32(minGasLimit))
		if err != nil {
			return fmt.Errorf("craft commitBatch tx failed:%v", err)
		}
	} else {
		versionedHashes := make([]common.Hash, 0)
		for _, commit := range batch.Sidecar.Commitments {
			versionedHashes = append(versionedHashes, kZGToVersionedHash(commit))
		}
		tx = types.NewTx(&types.BlobTx{
			ChainID:    uint256.MustFromBig(sr.chainId),
			Nonce:      nonce,
			GasTipCap:  uint256.MustFromBig(big.NewInt(tip.Int64())),
			GasFeeCap:  uint256.MustFromBig(big.NewInt(gasFeeCap.Int64())),
			Gas:        5000000,
			To:         sr.rollupAddr,
			Value:      uint256.NewInt(0),
			Data:       calldata,
			BlobFeeCap: uint256.NewInt(10e10),
			BlobHashes: versionedHashes,
			Sidecar: &types.BlobTxSidecar{
				Blobs:       batch.Sidecar.Blobs,
				Commitments: batch.Sidecar.Commitments,
				Proofs:      batch.Sidecar.Proofs,
			},
		})
	}

	newTx, err := UpdateGasLimit(tx)
	if err != nil {
		return fmt.Errorf("update gaslimit error:%v", err)
	}

	var newSignedTx *types.Transaction
	if tx.Type() == types.BlobTxType {
		newSignedTx, err = types.SignTx(newTx, types.NewLondonSignerWithEIP4844(sr.chainId), sr.privKey)
		if err != nil {
			return fmt.Errorf("sign tx error:%v", err)
		}
	} else {
		newSignedTx, err = opts.Signer(opts.From, newTx)
		if err != nil {
			return fmt.Errorf("sign tx error:%v", err)
		}
	}

	tx_send_time := time.Now().Unix()
	var receipt *types.Receipt
	err = sr.L1Client.SendTransaction(context.Background(), newSignedTx)
	if err != nil {
		log.Info("send tx error", "error", err.Error())

		// ErrReplaceUnderpriced,ErrAlreadyKnown
		if utils.ErrStringMatch(err, core.ErrReplaceUnderpriced) ||
			utils.ErrStringMatch(err, core.ErrAlreadyKnown) {
			receipt, newSignedTx, err = sr.replaceTx(newSignedTx)
			if err != nil {
				return fmt.Errorf("replace tx error:%v", err)
			} else {
				log.Info("replace tx success")
			}
		} else if utils.ErrStringMatch(err, core.ErrGasLimit) { // ErrGasLimit
			log.Error("tx exceeds block gas limit", "gas", newSignedTx.Gas(), "chunks_len", len(batch.Chunks))
			return fmt.Errorf("send tx error:%v", err.Error())
		} else if utils.ErrStringMatch(err, core.ErrNonceTooLow) { //ErrNonceTooLow
			return fmt.Errorf("send tx error:%v", err.Error())
		} else {
			return fmt.Errorf("send tx error:%v", err.Error())
		}

		// switch err.Error() {
		// case
		// 	core.ErrReplaceUnderpriced.Error(),
		// 	core.ErrAlreadyKnown.Error():

		// 	receipt, newSignedTx, err = sr.replaceTx(newSignedTx)
		// 	if err != nil {
		// 		return fmt.Errorf("replace tx error:%v", err)
		// 	} else {
		// 		log.Info("replace tx success")
		// 	}
		// case core.ErrGasLimit.Error():
		// 	log.Error("tx exceeds block gas limit", "gas", newSignedTx.Gas(), "chunks_len", len(batch.Chunks))
		// 	return fmt.Errorf("send tx error:%v", err.Error())
		// case core.ErrNonceTooLow.Error():
		// 	return fmt.Errorf("send tx error:%v", err.Error())
		// default:
		// 	return fmt.Errorf("send tx error:%v", err.Error())
		// }
	} else {
		log.Info("tx sent",
			// for business
			"batch_index", index,
			"chunks_len", len(batch.Chunks),
			// tx
			"tx_hash", newSignedTx.Hash().String(),
			"type", newSignedTx.Type(),
			"gas", newSignedTx.Gas(),
			"nonce", newSignedTx.Nonce(),
			"size", newSignedTx.Size(),
			"gas_price", newSignedTx.GasPrice().String(),
			"tip", newSignedTx.GasTipCap().String(),
			"fee_cap", newSignedTx.GasFeeCap().String(),
			"blob_fee_cap", newSignedTx.BlobGasFeeCap(),
			"blob_gas", newSignedTx.BlobGas(),
		)
	}

	// wait receipt
	if receipt == nil {
		receipt, err = sr.waitReceiptWithTimeout(sr.txTimout, newSignedTx.Hash())
		if err != nil {
			log.Error("wait receipt error", "error", err.Error())

			if utils.ErrStringMatch(err, errors.New("timeout")) {
				receipt, newSignedTx, err = sr.replaceTx(newSignedTx)
				if err != nil {
					return fmt.Errorf("replace tx error:%v", err)
				} else {
					log.Info("replace tx success")
				}
			} else {
				return fmt.Errorf("wait receipt error:%v", err)
			}
		}
	}

	tx_receipt_time := time.Now().Unix()
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("rollup tx failed",
			//block info
			"block_number", receipt.BlockNumber.Uint64(),
			// for business
			"batch_index", index,
			"chunks_len", len(batch.Chunks),
			// receipt info
			"gas_used", receipt.GasUsed,
			// tx info
			"tx_hash", newSignedTx.Hash().String(),
			"type", newSignedTx.Type(),
			"gas", newSignedTx.Gas(),
			"nonce", newSignedTx.Nonce(),
			"size", newSignedTx.Size(),
			"gas_price", newSignedTx.GasPrice().String(),
			"tip", newSignedTx.GasTipCap().String(),
			"fee_cap", newSignedTx.GasFeeCap().String(),
			"tx_included_time", tx_receipt_time-tx_send_time,
		)
		return fmt.Errorf("tx failed")
	} else {

		gasFee := newSignedTx.GasPrice().Uint64() * receipt.GasUsed
		gasFeeEther := new(big.Rat).SetFrac(big.NewInt(int64(gasFee)), big.NewInt(params.Ether))
		log.Info("rollup tx success",
			//block info
			"block_number", receipt.BlockNumber.Uint64(),
			// for business
			"batch_index", index,
			"chunks_len", len(batch.Chunks),
			// receipt info
			"gas_used", receipt.GasUsed,
			"tx_hash", newSignedTx.Hash().String(),
			// tx info
			"tx_hash", newSignedTx.Hash().String(),
			"type", newSignedTx.Type(),
			"gas", newSignedTx.Gas(),
			"nonce", newSignedTx.Nonce(),
			"size", newSignedTx.Size(),
			"gas_price", newSignedTx.GasPrice().String(),
			"tip", newSignedTx.GasTipCap().String(),
			"fee_cap", newSignedTx.GasFeeCap().String(),
			"gas_fee", gasFeeEther.FloatString(18),
			"tx_included_time", tx_receipt_time-tx_send_time,
		)
	}
	return nil
}

func (sr *SR) aggregateSignatures(blsSignatures []eth.RPCBatchSignature) (*bindings.IRollupBatchSignature, error) {
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

func (sr *SR) GetGasTipAndCap() (*big.Int, *big.Int, error) {
	tip, err := sr.L1Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, err
	}
	head, err := sr.L1Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, err
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
	return tip, gasFeeCap, nil
}

func (sr *SR) waitForReceipt(txHash common.Hash) (*types.Receipt, error) {
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

func (sr *SR) waitReceiptWithTimeout(time time.Duration, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(sr.ctx, time)
	defer cancel()
	return sr.waitReceiptWithCtx(ctx, txHash)
}

func (sr *SR) waitReceiptWithCtx(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
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
func (sr *SR) Init() error {
	// check whether the submitter is staked
	isSequencer, err := sr.Rollup.IsSequencer(
		&bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		crypto.PubkeyToAddress(sr.privKey.PublicKey),
	)

	if err != nil {
		return err
	}
	if !isSequencer {
		return fmt.Errorf("this account is not sequencer")
	}
	minDeposit, err := sr.Rollup.MINDEPOSIT(nil)
	if err != nil {
		return fmt.Errorf("query min deposit error:%v", err)
	}

	depositAmt, err := sr.Rollup.Deposits(nil, crypto.PubkeyToAddress(sr.privKey.PublicKey))
	if err != nil {
		return err
	}
	if depositAmt.Cmp(minDeposit) < 0 {
		// make sure the wallet balance is enough
		balance, err := sr.L1Client.BalanceAt(context.Background(), crypto.PubkeyToAddress(sr.privKey.PublicKey), nil)
		if err != nil {
			return err
		}
		log.Info("wallet info", "balance", new(big.Int).Div(balance, big.NewInt(params.Ether)).String())
		value := big.NewInt(0).Sub(minDeposit, depositAmt)
		if balance.Cmp(big.NewInt(0).Sub(minDeposit, depositAmt)) < 1 {
			return errors.New("balance not enough for staking")
		}
		// try to stake
		log.Info("submitter is not staked, try to stake")
		opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
		if err != nil {
			return err
		}

		opts.Value = value // 1 ether
		tx, err := sr.Rollup.Stake(opts)
		if err != nil {
			return err
		}

		receipt, err := bind.WaitMined(context.Background(), sr.L1Client, tx)
		if err != nil {
			return err
		}
		log.Info("stake success", "tx_hash", receipt.TxHash.String())
	}
	return nil
}

func (sr *SR) walletAddr() string {
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

func (sr *SR) replaceTx(tx *types.Transaction) (*types.Receipt, *types.Transaction, error) {
	if tx == nil {
		return nil, nil, errors.New("nil tx")
	}

	// for sign
	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return nil, nil, fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	// replaced tx info
	log.Info("replaced tx",
		"tx_hash", tx.Hash().String(),
		"gas_fee_cap", tx.GasFeeCap(),
		"gas_tip", tx.GasTipCap(),
		"gas", tx.Gas(),
		"nonce", tx.Nonce(),
	)

	_tip, _gasFeeCap, err := sr.GetGasTipAndCap()
	if err != nil {
		log.Error("get tip and cap", "err", err)
	}
	tip := _tip.Int64()
	gasFeeCap := _gasFeeCap.Int64()

	txTip := tx.GasTipCap().Int64()
	txGasFeeCap := tx.GasFeeCap().Int64()

	minTip := txTip * 12 / 10
	minGasFeeCap := txGasFeeCap * 12 / 10

	for {
		if (tip > minTip) && (gasFeeCap > minGasFeeCap) {
			break
		}
		if tip < minTip {
			tip *= 2
		}
		if gasFeeCap < minGasFeeCap {
			gasFeeCap *= 2
		}

	}

	// try 10 times
	for i := 0; i < 10; i++ {

		newTx := types.NewTx(&types.DynamicFeeTx{
			To:        tx.To(),
			Nonce:     tx.Nonce(),
			GasFeeCap: big.NewInt(gasFeeCap),
			GasTipCap: big.NewInt(tip),
			Gas:       tx.Gas(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
		log.Info("new tx info",
			"gas_tip", fmt.Sprintf("%d", tip), //todo: convert to gwei
			"gas_fee_cap", fmt.Sprintf("%d", gasFeeCap), //todo: convert to gwei
		)
		// sign tx
		opts.Nonce = big.NewInt(int64(newTx.Nonce()))
		newTx, err = opts.Signer(opts.From, newTx)
		if err != nil {
			return nil, nil, err
		}
		// send tx
		err = sr.L1Client.SendTransaction(context.Background(), newTx)
		if err != nil { // tx rejected
			log.Error("send replace tx", "err", err)
			// if not underprice return

			if !utils.ErrStringMatch(err, core.ErrReplaceUnderpriced) {
				return nil, nil, fmt.Errorf("send tx in replace error:%v", err.Error())
			}
		} else { // tx into mempool
			receipt, err := sr.waitReceiptWithTimeout(sr.txTimout, newTx.Hash())
			if err != nil {
				log.Error("wait receipt", "err", err)
				if err.Error() == "timeout" {
					log.Error("wait receipt timeout",
						"turns", i,
						"tip", tip,
						"gas_fee_cap", gasFeeCap,
					)
				} else {
					return nil, nil, err
				}
			} else {
				return receipt, newTx, nil
			}
		}
		// update tip & cap
		tip *= 2
		gasFeeCap *= 2

	}
	return nil, nil, errors.New("replace tx failed after try 10 times")
}
