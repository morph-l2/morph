package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/tx-submitter/iface"
	"github.com/morph-l2/tx-submitter/metrics"
	"github.com/morph-l2/tx-submitter/utils"

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

	for {
		if err := sr.rollup(); err != nil {
			if utils.IsRpcErr(err) {
				sr.metrics.IncRpcErrors()
			}
			time.Sleep(2 * time.Second)
			log.Error("rollup failed,wait for the next try", "error", err)
		}
		// call finalize
		if sr.Finalize {
			if err := sr.finalize(); err != nil {
				log.Error("finalize failed", "error", err)
			}
		} else {
			log.Info("finalize is disabled")
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
	if err := sr.L1Client.SendTransaction(context.Background(), newSignedTx); err != nil {
		if err.Error() != "already known" {
			return fmt.Errorf("SendTransaction error:%v", err.Error())
		}
	}
	// wait receipt
	receipt, err := sr.waitReceiptWithTimeOut(sr.txTimout, newSignedTx.Hash())
	if err != nil {
		return fmt.Errorf("wait receipt error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("finalize tx failed", "tx_hash", newSignedTx.Hash().String(), "finalize_cnt", finalizeCnt, "gas_used", receipt.GasUsed)
		return fmt.Errorf("tx failed")
	} else {
		log.Info("finalize tx success",
			"tx_hash", newSignedTx.Hash().String(),
			"gas_used", receipt.GasUsed,
			"gas_price", newSignedTx.GasPrice().String(),
			"gas_fee", new(big.Int).Mul(newSignedTx.GasPrice(), big.NewInt(int64(receipt.GasUsed))).String(),
			"last_finalized_before", lastFinalized.Uint64(),
			"last_commited_before", lastCommited.Uint64(),
			"finalize_cnt", finalizeCnt,
		)
	}
	return nil

}

func (sr *SR) rollup() error {

	if sr.PriorityRollup {
		log.Info("priority rollup is enabled")
	} else {
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
		return err
	}

	index := batchIndex.Uint64() + 1
	log.Info("batch info", "last_commit_batch", batchIndex.Uint64(), "batch_will_get", index)
	batch, err := GetRollupBatchByIndex(index, sr.L2Clients)
	if err != nil {
		return err
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

	var signedTx *types.Transaction

	opts, err := bind.NewKeyedTransactorWithChainID(sr.privKey, sr.chainId)
	if err != nil {
		return fmt.Errorf("new keyedTransaction with chain id error:%v", err)
	}

	opts.NoSend = true
	opts.Nonce = big.NewInt(int64(nonce))
	signedTx, err = sr.Rollup.CommitBatch(opts, rollupBatch, 1000)
	if err != nil {
		return fmt.Errorf("craft CommitBatch tx failed:%v", err)
	}
	if uint64(signedTx.Size()) > txMaxSize {
		return core.ErrOversizedData
	}

	newTx, err := UpdateGasLimit(signedTx)
	if err != nil {
		return fmt.Errorf("update gas limit error:%v", err)
	}
	newSignedTx, err := opts.Signer(opts.From, newTx)
	if err != nil {
		return fmt.Errorf("sign tx error:%v", err)
	}
	if err := sr.L1Client.SendTransaction(context.Background(), newSignedTx); err != nil {
		if err.Error() != "already known" {
			return fmt.Errorf("SendTransaction error:%v", err.Error())
		}
	}

	// wait receipt
	receipt, err := sr.waitReceiptWithTimeOut(sr.txTimout, newSignedTx.Hash())
	if err != nil {
		return fmt.Errorf("wait receipt error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Error("rollup tx failed", "tx_hash", newSignedTx.Hash().String(), "batch_index", index, "chunks_size", len(batch.Chunks))
		return fmt.Errorf("tx failed")
	} else {
		log.Info("rollup tx success",
			"tx_hash", newSignedTx.Hash().String(),
			"batch_index", index,
			"chunks_size", len(batch.Chunks),
			"gas_used", receipt.GasUsed,
			"gas_price", newSignedTx.GasPrice().String(),
			"gas_fee", new(big.Int).Mul(newSignedTx.GasPrice(), big.NewInt(int64(receipt.GasUsed))).String(),
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

func (sr *SR) waitReceiptWithTimeOut(time time.Duration, txHash common.Hash) (*types.Receipt, error) {
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
			GasPrice: tx.GasPrice(),
			Gas:      newGasLimit,
			To:       tx.To(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		})
	} else if tx.Type() == types.DynamicFeeTxType {
		newTx = types.NewTx(&types.DynamicFeeTx{
			Nonce:     tx.Nonce(),
			GasTipCap: tx.GasTipCap(),
			GasFeeCap: tx.GasFeeCap(),
			Gas:       newGasLimit,
			To:        tx.To(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
	} else {
		return nil, fmt.Errorf("unknown tx type:%v", tx.Type())
	}
	return newTx, nil
}
