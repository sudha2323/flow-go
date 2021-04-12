package flattener_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/complete/mtrie"
	"github.com/onflow/flow-go/ledger/complete/mtrie/flattener"
	"github.com/onflow/flow-go/module/metrics"
)

func TestForestStoreAndLoad(t *testing.T) {

	metricsCollector := &metrics.NoopCollector{}
	mForest, err := mtrie.NewForest(5, metricsCollector, nil)
	require.NoError(t, err)
	rootHash := mForest.GetEmptyRootHash()

	p1 := ledger.PathByUint8(1)
	v1 := ledger.LightPayload8('A', 'a')
	p2 := ledger.PathByUint8(2)
	v2 := ledger.LightPayload8('B', 'b')
	p3 := ledger.PathByUint8(130)
	v3 := ledger.LightPayload8('C', 'c')
	p4 := ledger.PathByUint8(131)
	v4 := ledger.LightPayload8('D', 'd')
	p5 := ledger.PathByUint8(132)
	v5 := ledger.LightPayload8('E', 'e')

	paths := []ledger.Path{p1, p2, p3, p4, p5}
	payloads := []*ledger.Payload{v1, v2, v3, v4, v5}

	update := &ledger.TrieUpdate{RootHash: rootHash, Paths: paths, Payloads: payloads}
	rootHash, err = mForest.Update(update)
	require.NoError(t, err)

	p6 := ledger.PathByUint8(133)
	v6 := ledger.LightPayload8('F', 'f')
	update = &ledger.TrieUpdate{RootHash: rootHash, Paths: []ledger.Path{p6}, Payloads: []*ledger.Payload{v6}}
	rootHash, err = mForest.Update(update)
	require.NoError(t, err)

	forestSequencing, err := flattener.FlattenForest(mForest)
	require.NoError(t, err)

	newForest, err := mtrie.NewForest(5, metricsCollector, nil)
	require.NoError(t, err)

	//forests are different
	assert.NotEqual(t, mForest, newForest)

	rebuiltTries, err := flattener.RebuildTries(forestSequencing)
	require.NoError(t, err)
	err = newForest.AddTries(rebuiltTries)
	require.NoError(t, err)

	//forests are the same now
	assert.Equal(t, mForest, newForest)

	read := &ledger.TrieRead{RootHash: rootHash, Paths: paths}
	retPayloads, err := mForest.Read(read)
	require.NoError(t, err)
	newRetPayloads, err := newForest.Read(read)
	require.NoError(t, err)
	for i := range paths {
		require.True(t, retPayloads[i].Equals(newRetPayloads[i]))
	}
}
