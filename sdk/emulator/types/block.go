package types

import (
	"time"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/dapperlabs/flow-go/pkg/crypto"
)

type Block struct {
	Height            uint64
	Timestamp         time.Time
	PreviousBlockHash crypto.Hash
	TransactionHashes []crypto.Hash
}

func (b *Block) Hash() crypto.Hash {
	hasher, _ := crypto.NewHasher(crypto.SHA3_256)

	d, _ := rlp.EncodeToBytes([]interface{}{
		b.Height,
		b.PreviousBlockHash,
	})

	return hasher.ComputeHash(d)
}

func GenesisBlock() *Block {
	return &Block{
		Height:            0,
		Timestamp:         time.Now(),
		PreviousBlockHash: nil,
		TransactionHashes: make([]crypto.Hash, 0),
	}
}
