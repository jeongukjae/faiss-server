package faiss

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadIndex(t *testing.T) {
	idx, err := LoadIndex("../testdata/random-index.faiss")
	assert.Nil(t, err)
	assert.NotNil(t, idx)
}

func TestGetMetadata(t *testing.T) {
	idx, err := LoadIndex("../testdata/random-index.faiss")
	assert.Nil(t, err)
	assert.NotNil(t, idx)

	metadata := idx.GetMetadata()
	assert.Equal(t, metadata.Dimension, 64)
	assert.Equal(t, metadata.Ntotal, 1000)
	assert.Equal(t, metadata.MetricType, 1)
}
