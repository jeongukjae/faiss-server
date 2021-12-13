package faiss

/*
#include "faiss/wrapper.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type FaissIndex struct {
	Index *C.FaissIndex

	Path       string
	Dimension  int32
	MetricType int32
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

	dimension := int32(C.faiss_Index_d(index))
	metricType := int32(C.faiss_Index_metric_type(index))
	return &FaissIndex{
		Index:      index,
		Path:       path,
		Dimension:  dimension,
		MetricType: metricType,
	}, nil
}

func (index *FaissIndex) Free() {
	C.faiss_Index_free(index.Index)
}

func (index *FaissIndex) GetNtotal() int64 {
	Ntotal := int64(C.faiss_Index_ntotal(index.Index))
	return Ntotal
}

func (index *FaissIndex) Search(numVectors int32, topK int32, vectors []float32) *SearchResult {
	cSearchResult := C.searchFaiss(index.Index, C.int(numVectors), C.int(topK), (*C.float)(&vectors[0]))
	numResults := topK * numVectors

	cIds := unsafe.Pointer(cSearchResult.ids)
	cDistances := unsafe.Pointer(cSearchResult.distances)

	cIdsArray := (*[1 << 30]C.int64_t)(cIds)
	cDistancesArray := (*[1 << 30]C.float)(cDistances)

	defer C.free(cIds)
	defer C.free(cDistances)

	ids := make([]int64, numResults)
	distances := make([]float32, numResults)

	for i := int32(0); i < numResults; i++ {
		ids[i] = int64(cIdsArray[i])
		distances[i] = float32(cDistancesArray[i])
	}

	return &SearchResult{
		Ids:       ids,
		Distances: distances,
	}
}

func (index *FaissIndex) AddVectors(numVectors int32, vectors []float32) ([]int64, error) {
	ids := make([]int64, numVectors)
	for i := range ids {
		ids[i] = index.GetNtotal() + int64(i)
	}
	code := C.faiss_Index_add(index.Index, C.idx_t(numVectors), (*C.float)(&vectors[0]))
	if int(code) != 0 {
		return nil, errors.New("Cannot add vectors")
	}
	return ids, nil
}

func (index *FaissIndex) AddVectorsWithIds(numVectors int32, vectors []float32, ids []int64) error {
	code := C.faiss_Index_add_with_ids(index.Index, C.idx_t(numVectors), (*C.float)(&vectors[0]), (*C.idx_t)(&ids[0]))
	if int(code) != 0 {
		return errors.New("Cannot add vectors")
	}
	return nil
}

func (index *FaissIndex) RemoveVectors(ids []int64) (int32, error) {
	numRemovedOrCode := int32(C.removeVectors(index.Index, C.int(len(ids)), (*C.idx_t)(&ids[0])))
	if numRemovedOrCode == -1 {
		return -1, errors.New("Cannot remove vectors")
	}
	return numRemovedOrCode, nil
}
