package validator

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind/backends"
	"github.com/scroll-tech/go-ethereum/core"
	"github.com/scroll-tech/go-ethereum/core/rawdb"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethdb"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestValidator_ChallengeState(t *testing.T) {
	key, _ := crypto.GenerateKey()
	sim, _ := newSimulatedBackend(key)
	opts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)
	addr, _, rollup, err := bindings.DeployRollup(opts, sim, 1337)
	require.NoError(t, err)
	sim.Commit()
	v := Validator{
		cli:        sim,
		privateKey: key,
		l1ChainID:  big.NewInt(1),
		contract:   rollup,
	}
	err = v.ChallengeState(10)
	log.Info("addr:", addr)
	require.EqualError(t, err, "execution reverted: Batch not exist")
}

func newSimulatedBackend(key *ecdsa.PrivateKey) (*backends.SimulatedBackend, ethdb.Database) {
	var gasLimit uint64 = 9_000_000
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}
	db := rawdb.NewMemoryDatabase()
	sim := backends.NewSimulatedBackendWithDatabase(db, genAlloc, gasLimit)
	return sim, db
}
