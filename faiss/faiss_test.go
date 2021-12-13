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

	assert.Equal(t, idx.Dimension, int32(8))
	assert.Equal(t, idx.GetNtotal(), int32(1000))
	assert.Equal(t, idx.MetricType, int32(1))
}

func TestSearch(t *testing.T) {
	idx, err := LoadIndex("../testdata/random-index.faiss")
	assert.Nil(t, err)
	assert.NotNil(t, idx)

	queryVectors := []float32{
		0.497664, 0.817838, 0.612112, 0.77136, 0.86067, 0.150637, 0.198519, 0.815163,
		0.159815, 0.116138, 0.0129075, 0.486833, 0.331015, 0.80264, 0.0982519, 0.0559934,
		0.444663, 0.0221439, 0.290729, 0.246394, 0.738287, 0.889226, 0.987139, 0.117443,
		0.396782, 0.45273, 0.538148, 0.790622, 0.465836, 0.435332, 0.569479, 0.969259,
		0.0445562, 0.54812, 0.462577, 0.376472, 0.327912, 0.813529, 0.646552, 0.0474265,
	}
	results := idx.Search(5, 4, queryVectors)

	assert.ElementsMatch(t, []int64{
		0, 372, 74, 425,
		1, 27, 221, 538,
		2, 170, 40, 60,
		3, 39, 79, 627,
		4, 198, 67, 282,
	}, results.Ids)
	ElementsAllClose(t, []float32{
		0, 0.220977, 0.23984, 0.298714,
		0, 0.379492, 0.382708, 0.503856,
		0, 0.341152, 0.444937, 0.456188,
		0, 0.108953, 0.302287, 0.304529,
		0, 0.371397, 0.377862, 0.431174,
	}, results.Distances, 1e-8, 1e-5)
}

func ElementsAllClose(t *testing.T, listA, listB []float32, atol float32, rtol float32, msgAndArgs ...interface{}) bool {
	if len(listA) != len(listB) {
		t.Errorf("ElementsAllClose Failed \nlistA: %s\nlistB: %s\natol: %f\nrtol: %f", listA, listB, atol, rtol)
	}

	for i, _ := range listA {
		if Abs(listA[i]-listB[i]) > (atol + rtol*Abs(listB[i])) {
			t.Errorf("ElementsAllClose Failed at index %d\nlistA: %s\nlistB: %s\natol: %f\nrtol: %f", i, listA, listB, atol, rtol)
		}
	}

	return true
}

func Abs(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}
