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
