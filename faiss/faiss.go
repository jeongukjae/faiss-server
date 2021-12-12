package faiss

/*
#include "faiss/faiss_c_wrapper.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type FaissIndex struct {
	Index *C.FaissIndex

	Path string
}

type FaissMetadata struct {
	Dimension  int32
	Ntotal     int32
	MetricType int32
}

type SearchResult struct {
	Ids       []int64
	Distances []float32
}

func LoadIndex(path string) (*FaissIndex, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	index := C.loadIndex(cPath)
	if index == nil {
		return nil, errors.New("Cannot create index")
	}
	return &FaissIndex{
		Index: index,
		Path:  path,
	}, nil
}

func (index *FaissIndex) Free() {
	C.free(unsafe.Pointer(index.Index))
}

func (index *FaissIndex) GetMetadata() *FaissMetadata {
	cMetadata := C.getMetadata(index.Index)
	defer C.free(unsafe.Pointer(cMetadata))

	metadata := FaissMetadata{
		Dimension:  int32(cMetadata.dimension),
		Ntotal:     int32(cMetadata.ntotal),
		MetricType: int32(cMetadata.metric_type),
	}

	return &metadata
}

func (index *FaissIndex) Search(numVectors int, vectors []float32, topK int) *SearchResult {
	cSearchResult := C.searchFaiss(index.Index, C.int(numVectors), (*C.float)(&vectors[0]), C.int(topK))
	numResults := topK * numVectors

	cIds := unsafe.Pointer(cSearchResult.ids)
	cDistances := unsafe.Pointer(cSearchResult.distances)

	cIdsArray := (*[1 << 30]C.int64_t)(cIds)
	cDistancesArray := (*[1 << 30]C.float)(cDistances)

	defer C.free(cIds)
	defer C.free(cDistances)

	ids := make([]int64, numResults)
	distances := make([]float32, numResults)

	for i := 0; i < numResults; i++ {
		ids[i] = int64(cIdsArray[i])
		distances[i] = float32(cDistancesArray[i])
	}

	return &SearchResult{
		Ids:       ids,
		Distances: distances,
	}
}
