package oracle

import (
	"context"
	"fmt"
	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/bindings/predeploys"
	"github.com/morph-l2/morph/oracle/config"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"strings"
	"testing"
	"time"
)

func testNewOracleClient(t *testing.T) *Oracle {
	l1Client, err := ethclient.Dial("http://localhost:9545")
	require.NoError(t, err)
	var secret [32]byte
	jwtSecret := common.FromHex(strings.TrimSpace("688f5d737bad920bdfb2fc2f488d6b6209eebda1dae949a8de91398d932c517a"))
	require.True(t, len(jwtSecret) == 32)
	copy(secret[:], jwtSecret)
	l2Client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	//logger := tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))
	l2Staking, err := bindings.NewL2Staking(predeploys.L2StakingAddr, l2Client)
	if err != nil {
		panic(err)
	}
	record, err := bindings.NewRecord(predeploys.RecordAddr, l2Client)
	if err != nil {
		panic(err)
	}
	cfg := &config.Config{
		StartBlock: 1,
	}
	return &Oracle{
		l1Client: l1Client,
		l2Client: l2Client,
		//rollup:      rollup,
		l2Staking:   l2Staking,
		record:      record,
		cfg:         cfg,
		rewardEpoch: defaultRewardEpoch,
		ctx:         context.Background(),
	}
}

func TestSetStartBlock(t *testing.T) {
	o := testNewOracleClient(t)
	rewardStartTime, err := o.l2Staking.RewardStartTime(nil)
	require.NoError(t, err)
	fmt.Println("rewardStartTime:", rewardStartTime.Uint64())

	o.setStartBlock()
}

func TestGetStakers(t *testing.T) {
	o := testNewOracleClient(t)
	stakers, err := o.l2Staking.GetStakers(nil)
	require.NoError(t, err)
	for _, staker := range stakers {
		fmt.Println("staker addr:", staker.Addr)
		commission, err := o.l2Staking.Commissions(nil, staker.Addr)
		require.NoError(t, err)
		fmt.Println("commission:", commission)
	}
}

func TestQueryRewardStartTime(t *testing.T) {
	o := testNewOracleClient(t)
	rewardStartTime, err := o.l2Staking.RewardStartTime(nil)
	require.NoError(t, err)
	fmt.Println("rewardStartTime:", rewardStartTime.Uint64())
	fmt.Println(time.Now().Unix())
	fmt.Println(rewardStartTime.Uint64())

}

func TestSetRewardStart(t *testing.T) {
	o := testNewOracleClient(t)
	chainID, err := o.l2Client.ChainID(o.ctx)
	if err != nil {

	}
	hex := strings.TrimPrefix("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		return
	}
	opts.NoSend = true
	opts.Nonce = big.NewInt(int64(nonce))
	recordOwner, err := o.record.Owner(nil)
	require.NoError(t, err)
	fmt.Println(recordOwner.Hex())
	owner, err := o.l2Staking.Owner(nil)
	require.NoError(t, err)
	fmt.Println(owner.Hex())
	tx1, err := o.l2Staking.UpdateRewardStartTime(opts, big.NewInt(int64(time.Now().Add(30*time.Minute).Unix())))
	require.NoError(t, err)
	signedTx1, err := opts.Signer(opts.From, tx1)
	require.NoError(t, err)
	fmt.Println(signedTx1)
	fmt.Println("Tx:", tx1.Hash())
	require.NoError(t, err)
	tx, err := o.l2Staking.StartReward(opts)
	require.NoError(t, err)
	fmt.Println("ops from:", opts.From)
	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return
	}
	fmt.Println(signedTx)
	//err = o.l2Client.SendTransaction(o.ctx, signedTx)
	//if err != nil {
	//
	//}
	//receipt, err := o.l2Client.TransactionReceipt(o.ctx, tx.Hash())
	//if receipt.Status != types.ReceiptStatusSuccessful {
	//	fmt.Println("set stark block failed")
	//}
	require.NoError(t, err)
	fmt.Println("Tx:", tx.Hash())
}
