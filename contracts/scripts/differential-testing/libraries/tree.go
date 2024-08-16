package libraries

import (
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
)

type SMT struct {
	Height     uint64
	zeroHashes []common.Hash
	tree       [][]common.Hash
	LeafNum    uint64
	dirty      bool
}

func NewSMT(height uint64) *SMT {
	return &SMT{
		Height:     height,
		zeroHashes: generateZeroHashes(),
		tree:       generateZeroTrees(),
		LeafNum:    0,
		dirty:      false,
	}
}

func (t *SMT) Get(index uint64) common.Hash {
	return t.tree[0][index]
}

func (t *SMT) Add(leaf common.Hash) {
	t.dirty = true
	allRoll := t.tree[0]
	allRoll = append(allRoll, leaf)
	t.tree[0] = allRoll
	t.LeafNum++
}

func (t *SMT) calcBranches() {
	for i := uint64(0); i < t.Height; i++ {
		parent := t.tree[i+1]
		child := t.tree[i]
		for j := 0; j < len(child); j += 2 {
			leftNode := child[j]
			rightNode := leftNode
			if j+1 < len(child) {
				rightNode = child[j+1]
			} else {
				rightNode = t.zeroHashes[i]
			}
			if j+1 < len(child) {
				rightNode = child[j+1]
			} else {
				rightNode = t.zeroHashes[i]
			}
			if len(parent) <= j/2 {
				allRoll := t.tree[i+1]
				allRoll = append(allRoll, crypto.Keccak256Hash(leftNode[:], rightNode[:]))
				t.tree[i+1] = allRoll
			} else {
				parent[j/2] = crypto.Keccak256Hash(leftNode.Bytes(), rightNode.Bytes())
			}
		}
	}
	t.dirty = false
}

func (t *SMT) GetRoot() common.Hash {
	if t.tree[0] == nil {
		return crypto.Keccak256Hash(t.zeroHashes[t.Height-1].Bytes(), t.zeroHashes[t.Height-1].Bytes())
	}
	if t.dirty {
		t.calcBranches()
	}

	return t.tree[t.Height][0]
}

func (t *SMT) GetProofTreeByIndex(index uint64) []common.Hash {
	if t.dirty {
		t.calcBranches()
	}
	proof := make([]common.Hash, 0)
	currentIndex := index
	for i := uint64(0); i < t.Height; i++ {
		if currentIndex%2 == 1 {
			currentIndex = currentIndex - 1
		} else {
			currentIndex = currentIndex + 1
		}
		if currentIndex < uint64(len(t.tree[i])) {
			proof = append(proof, t.tree[i][currentIndex])
		} else {
			proof = append(proof, t.zeroHashes[i])
		}
		currentIndex = currentIndex / 2
	}
	return proof
}

func (t *SMT) GetProofTreeByValue(value common.Hash) []common.Hash {
	index := -1
	for i, leaf := range t.tree[0] {
		if leaf == value {
			index = i
			break
		}
	}
	if index == -1 {
		return []common.Hash{}
	}
	return t.GetProofTreeByIndex(uint64(index))
}

func generateZeroTrees() [][]common.Hash {
	var nilTress [33][]common.Hash
	return nilTress[:]
}

func generateZeroHashes() []common.Hash {
	var zeroHashes [32]common.Hash

	for i := uint64(1); i < 32; i++ {
		hash := crypto.Keccak256Hash(zeroHashes[i-1].Bytes(), zeroHashes[i-1].Bytes())
		zeroHashes[i] = hash
	}
	return zeroHashes[:]
}

func VerifyMerkleProof(leaf common.Hash, smtProof []common.Hash, index int, root common.Hash) bool {
	value := leaf
	for i := 0; i < len(smtProof); i++ {
		if (index/(1<<i))%2 != 0 {
			value = crypto.Keccak256Hash(smtProof[i].Bytes(), value.Bytes())
		} else {
			value = crypto.Keccak256Hash(value.Bytes(), smtProof[i].Bytes())
		}
	}

	return value == root
}
