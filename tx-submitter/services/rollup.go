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

	secondInterval time.Duration
	txTimout       time.Duration
	Finalize       bool
	MaxFinalizeNum uint64
	PriorityRollup bool

	// tx cfg
	txFeeLimit uint64

	pendingTxs *PendingTxs

	rollupMu sync.Mutex
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

		secondInterval: cfg.PollInterval,
		txTimout:       cfg.TxTimeout,
		Finalize:       cfg.Finalize,
		MaxFinalizeNum: cfg.MaxFinalizeNum,
		PriorityRollup: cfg.PriorityRollup,

		txFeeLimit: cfg.TxFeeLimit,
	}
}

func (sr *Rollup) Start() {

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

	// init pendingtxs
	sr.pendingTxs = NewPendingTxs(sr.abi.Methods["commitBatch"].ID, sr.abi.Methods["finalizeBatch"].ID)
	txs, _, err := utils.ParseL1Mempool(sr.l1RpcClient, crypto.PubkeyToAddress(sr.privKey.PublicKey))
	if err != nil {
		log.Error("parse l1 mempool error", "error", err)
	}
	for _, tx := range txs {
		sr.pendingTxs.Add(tx)
	}

	var rollupMu sync.Mutex
	go utils.Loop(sr.ctx, time.Second*5, func() {
		rollupMu.Lock()
		defer rollupMu.Unlock()
		if err := sr.rollup(); err != nil {
			if utils.IsRpcErr(err) {
				sr.metrics.IncRpcErrors()
			}
			log.Error("rollup failed,wait for the next try", "error", err)
		}
	})

	if sr.Finalize {
		go utils.Loop(sr.ctx, time.Second*20, func() {
			if err := sr.finalize(); err != nil {
				log.Error("finalize failed", "error", err)
			}
		})
	}

	go utils.Loop(sr.ctx, time.Second*5, func() {
		if err := sr.ProcessTx(); err != nil {
			log.Error("process tx err", "error", err)
		}
	})

}

func (sr *Rollup) ProcessTx() error {

	// case 1: in mempool
	//          -> check timeout
	// case 2: no in mempool
	// case 2.1: disgard
	// case 2.2: tx included -> success
	// case 2.3: tx included -> failed
	//          -> log tx info

	// get all local txs
	txRecords := sr.pendingTxs.GetAll()
	if len(txRecords) == 0 {
		return nil
	}

	ptxs, _, err := utils.ParseL1Mempool(sr.l1RpcClient, crypto.PubkeyToAddress(sr.privKey.PublicKey))
	if err != nil {
		return fmt.Errorf("parse l1 mempool error:%w", err)
	}

	pendingtxMap := make(map[common.Hash]types.Transaction)
	for _, tx := range ptxs {
		pendingtxMap[tx.Hash()] = tx
	}

	// query tx status
	for _, txRecord := range txRecords {
		_, ok := pendingtxMap[txRecord.tx.Hash()]
		rtx := txRecord.tx

		// exist in mempool
		if ok {
			if txRecord.sendTime+uint64(sr.txTimout.Seconds()) < uint64(time.Now().Unix()) {
				_, err := sr.replaceTx(&rtx)
				if err != nil {
					log.Error("resubmit tx", "error", err)
					return fmt.Errorf("resubmit tx error:%w", err)
				} else {
					sr.pendingTxs.Remove(rtx.Hash())
				}
			}
		} else { // not in mempool
			receipt, err := sr.L1Client.TransactionReceipt(context.Background(), rtx.Hash())
			if err != nil {
				log.Error("query tx receipt error", "error", err)
				if utils.ErrStringMatch(err, ethereum.NotFound) {
					txRecord.queryTimes++
					if txRecord.queryTimes >= 5 {
						sr.pendingTxs.Remove(rtx.Hash())
					}
				} else {
					return fmt.Errorf("query tx receipt error:%w", err)
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
				sr.pendingTxs.Remove(rtx.Hash())
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

	target := new(big.Int).Add(lastFinalized, big.NewInt(1))
	if target.Cmp(lastCommited) > 0 {
		log.Info("no need to finalize", "last_finalized", lastFinalized.Uint64(), "last_committed", lastCommited.Uint64())
		return nil
	}
	// in challange window
	inWindow, err := sr.Rollup.BatchInsideChallengeWindow(nil, target)
	if err != nil {
		return fmt.Errorf("get batch inside challenge window error:%v", err)
	}
	if inWindow {
		log.Info("batch inside challenge window, wait for the next turn")
		return nil
	}
	// finalize
	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}
	opts.NoSend = true
	tx, err := sr.Rollup.FinalizeBatch(opts, target)
	if err != nil {
		return fmt.Errorf("craft FinalizeBatch tx failed:%v", err)
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

	log.Info("finalize tx info",
		"batch_index", target,
		"hash", newSignedTx.Hash().String(),
		"type", newSignedTx.Type(),
		"nonce", newSignedTx.Nonce(),
		"gas", newSignedTx.Gas(),
		"tip", newSignedTx.GasTipCap().String(),
		"fee_cap", newSignedTx.GasFeeCap().String(),
		"size", newSignedTx.Size(),
	)

	err = sr.SendTx(newSignedTx)
	if err != nil {
		log.Error("send tx to mempool", "error", err.Error())
		return fmt.Errorf("send tx error:%v", err.Error())
	} else {
		log.Info("finalzie tx sent")
	}

	return nil

}

func (sr *Rollup) rollup() error {

	if !sr.PriorityRollup {
		// is the turn of the submitter
		currentSubmitter, err := sr.getCurrentSubmitter()
		if err != nil {
			return fmt.Errorf("get next submitter error:%v", err)
		}
		log.Info("rotator info", "turn", currentSubmitter.Hex(), "cur", sr.walletAddr())

		if currentSubmitter.Hex() == sr.walletAddr() {
			log.Info("start to rollup")
		} else {
			log.Info("wait for my turn")
			return nil
		}
	}

	if len(sr.pendingTxs.txinfos) > 5 {
		log.Info("too many txs in mempool, wait")
		return nil
	}

	var nonce uint64
	var batchIndex uint64
	var err error

	if sr.pendingTxs.pindex != 0 {
		batchIndex = sr.pendingTxs.pindex + 1
	} else {
		batchIndexBig, err := sr.Rollup.LastCommittedBatchIndex(nil)
		if err != nil {
			return fmt.Errorf("get last committed batch index error:%v", err)
		}
		batchIndex = batchIndexBig.Uint64() + 1
	}

	log.Info("batch info", "last_commit_batch", batchIndex-1, "batch_will_get", batchIndex)
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
	// calldata encode
	calldata, err := sr.abi.Pack("commitBatch", rollupBatch, sequencerVersion, []common.Address{}, signature.Signature)
	if err != nil {
		return fmt.Errorf("pack calldata error:%v", err)
	}

	gas, err := sr.L1Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      opts.From,
		To:        &sr.rollupAddr,
		GasFeeCap: gasFeeCap,
		GasTipCap: tip,
		Data:      calldata,
	})
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %w", err)
	}

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

	// add a buffer to the gas limit
	newTx, err := UpdateGasLimit(tx)
	if err != nil {
		return fmt.Errorf("update gaslimit error:%v", err)
	}

	opts.Nonce = big.NewInt(int64(nonce))
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

	log.Info("rollup tx info",
		"batch_index", batchIndex,
		"hash", newSignedTx.Hash().String(),
		"type", newSignedTx.Type(),
		"nonce", newSignedTx.Nonce(),
		"gas", newSignedTx.Gas(),
		"tip", newSignedTx.GasTipCap().String(),
		"fee_cap", newSignedTx.GasFeeCap().String(),
		"blob_fee_cap", newSignedTx.BlobGasFeeCap(),
		"blob_gas", newSignedTx.BlobGas(),
		"size", newSignedTx.Size(),
		"blob_len", len(newSignedTx.BlobHashes()),
	)

	err = sr.SendTx(newSignedTx)
	if err != nil {
		log.Error("send tx to mempool", "error", err.Error())
		return fmt.Errorf("send tx error:%v", err.Error())
	} else {
		log.Info("rollup tx send to mempool succuess", "hash", newSignedTx.Hash().String())
	}

	return nil
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

	err := sendTx(r.L1Client, r.txFeeLimit, tx)
	if err != nil {
		return err
	}

	// after send tx
	// add to pending txs
	r.pendingTxs.Add(*tx)

	return nil

}

// send tx to l1 with business logic check
func sendTx(client iface.Client, txFeeLimit uint64, tx *types.Transaction) error {
	// fee limit
	if txFeeLimit > 0 {
		var fee uint64
		// calc tx gas fee
		if tx.Type() == types.BlobTxType {
			// blob fee
			fee = tx.BlobGasFeeCap().Uint64() * tx.BlobGas()
			// tx fee
			fee += tx.GasPrice().Uint64() * tx.Gas()
		} else {
			fee = tx.GasPrice().Uint64() * tx.Gas()
		}

		if fee > txFeeLimit {
			return fmt.Errorf("%v:limit=%v,but got=%v", utils.ErrExceedFeeLimit, txFeeLimit, fee)
		}
	}

	return client.SendTransaction(context.Background(), tx)
}

func (sr *Rollup) replaceTx(tx *types.Transaction) (*types.Transaction, error) {
	if tx == nil {
		return nil, errors.New("nil tx")
	}

	// for sign
	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return nil, fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	// replaced tx info
	log.Info("replaced tx",
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

	// bump tip & feeCap
	bumpedFeeCap := calcThresholdValue(tx.GasFeeCap(), tx.Type() == types.BlobTxType)
	bumpedTip := calcThresholdValue(tx.GasTipCap(), tx.Type() == types.BlobTxType)

	// if bumpedTip > tip
	if bumpedTip.Cmp(tip) > 0 {
		tip = bumpedTip
		gasFeeCap = bumpedFeeCap
	}
	if tx.Type() == types.BlobTxType {
		bumpedBlobFeeCap := calcBlobFeeCap(tx.BlobGasFeeCap())
		if bumpedBlobFeeCap.Cmp(blobFeeCap) > 0 {
			blobFeeCap = bumpedBlobFeeCap
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

func (r *Rollup) getCurrentSubmitter() (*common.Address, error) {

	for _, l2Submitter := range r.L2Submitters {
		current, _, _, err := l2Submitter.GetCurrentSubmitter(nil)
		if err != nil {
			log.Warn("get current submitter error", "error", err)
			continue
		}
		return &current, nil

	}

	return nil, errors.New("failed to get current submitter")
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
