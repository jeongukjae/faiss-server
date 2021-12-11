package faiss

/*
#include "faiss_wrapper.h"
*/
import "C"
import "unsafe"

func LoadIndex(path string) (*C.FaissIndex, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	index := C.loadIndex(cPath)
	return index, nil
}
