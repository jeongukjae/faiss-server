package faiss

/*
#include <stdlib.h>
#include "faiss/c_api/index_c.h"
#include "faiss/c_api/index_io_c.h"

FaissIndex* loadIndex(const char* path) {
    FaissIndex* index = NULL;
    if (faiss_read_index_fname(path, FAISS_IO_FLAG_READ_ONLY, &index)) {
        return NULL;
    }
    return index;
}
*/
import "C"
import "unsafe"

func LoadIndex(path string) (*C.FaissIndex, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	index := C.loadIndex(cPath)
	return index, nil
}
