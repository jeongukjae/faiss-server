#include "faiss_c_wrapper.h"

FaissIndex* loadIndex(const char* path) {
    FaissIndex* index = NULL;
    if (faiss_read_index_fname(path, FAISS_IO_FLAG_READ_ONLY, &index)) {
        return NULL;
    }
    return index;
}

FaissMetadata* getMetadata(FaissIndex* index) {
    FaissMetadata *metadata;
    metadata = malloc(sizeof(FaissMetadata));

    metadata->dimension = faiss_Index_d(index);
    metadata->ntotal = faiss_Index_ntotal(index);
    metadata->metric_type = faiss_Index_metric_type(index);

    return metadata;
}
