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
