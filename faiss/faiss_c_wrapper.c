#include "faiss_c_wrapper.h"

FaissIndex* loadIndex(const char* path) {
    FaissIndex* index = NULL;
    if (faiss_read_index_fname(path, FAISS_IO_FLAG_READ_ONLY, &index)) {
        return NULL;
    }
    return index;
}

FaissMetadata* getMetadata(const FaissIndex* index) {
    FaissMetadata *metadata;
    metadata = malloc(sizeof(FaissMetadata));

    metadata->dimension = faiss_Index_d(index);
    metadata->ntotal = faiss_Index_ntotal(index);
    metadata->metric_type = faiss_Index_metric_type(index);

    return metadata;
}

SearchResults searchFaiss(const FaissIndex* index, int numVectors, const float* vectors, int topK) {
    idx_t* ids = malloc(sizeof(idx_t) * topK * numVectors);
    float* distances = malloc(sizeof(float) * topK * numVectors);

    int result = faiss_Index_search(index, numVectors, vectors, topK, distances, ids);

    SearchResults searchResult = {
        ids,
        distances,
        result,
    };
    return searchResult;
}
