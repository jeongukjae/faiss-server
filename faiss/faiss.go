package faiss

/*
#include "faiss/faiss_c_wrapper.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func LoadIndex(path string) (*C.FaissIndex, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	index := C.loadIndex(cPath)
	if index == nil {
		return nil, errors.New("Cannot create index")
	}
	return index, nil
}
