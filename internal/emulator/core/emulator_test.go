package core_test

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/dapperlabs/bamboo-node/internal/emulator/core"
	"github.com/dapperlabs/bamboo-node/pkg/crypto"
	"github.com/dapperlabs/bamboo-node/pkg/types"
)

const sampleScript = `
	fun main() {
		const controller = [1]
		const owner = [2]
		const key = [3]
		const value = getValue(controller, owner, key)
		setValue(controller, owner, key, value + 2)
	}
`

func TestSubmitTransaction(t *testing.T) {
	RegisterTestingT(t)

	b := core.NewEmulatedBlockchain()

	txA := &types.SignedTransaction{
		Script:         []byte(sampleScript),
		Nonce:          1,
		ComputeLimit:   10,
		Timestamp:      time.Now(),
		PayerSignature: crypto.Signature{},
	}

	err := b.SubmitTransaction(txA)
	Expect(err).ToNot(HaveOccurred())

	txB := b.GetTransaction(txA.Hash())
	Expect(txB.Status).To(Equal(types.TransactionSealed))
}

func TestSubmitDuplicateTransaction(t *testing.T) {
	RegisterTestingT(t)

	b := core.NewEmulatedBlockchain()

	txA := &types.SignedTransaction{
		Script:         []byte(sampleScript),
		Nonce:          1,
		ComputeLimit:   10,
		Timestamp:      time.Now(),
		PayerSignature: crypto.Signature{},
	}

	err := b.SubmitTransaction(txA)
	Expect(err).ToNot(HaveOccurred())

	err = b.SubmitTransaction(txA)
	Expect(err).To(MatchError(&core.ErrDuplicateTransaction{TxHash: txA.Hash()}))
}

func TestSubmitRevertedTransaction(t *testing.T) {
	RegisterTestingT(t)

	b := core.NewEmulatedBlockchain()

	txA := &types.SignedTransaction{
		Script:         []byte("invalid script"),
		Nonce:          1,
		ComputeLimit:   10,
		Timestamp:      time.Now(),
		PayerSignature: crypto.Signature{},
	}

	err := b.SubmitTransaction(txA)
	Expect(err).To(HaveOccurred())

	txB := b.GetTransaction(txA.Hash())
	Expect(txB.Status).To(Equal(types.TransactionReverted))
}
