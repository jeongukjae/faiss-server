#include "faiss_wrapper.h"

FaissIndex* loadIndex(const char* path) {
    FaissIndex* index = NULL;
    if (faiss_read_index_fname(path, FAISS_IO_FLAG_READ_ONLY, &index)) {
        return NULL;
    }
    return index;
}
