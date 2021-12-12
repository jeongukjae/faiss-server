#ifndef __FAISS_WRAPPER_H__
#define __FAISS_WRAPPER_H__

#include <stdlib.h>
#include "faiss/c_api/index_c.h"
#include "faiss/c_api/index_io_c.h"

typedef struct FaissMetadata {
    int dimension;
    int ntotal;
    FaissMetricType metric_type;
} FaissMetadata;

typedef struct SearchResults {
    int64_t* ids;
    float* distances;
    int isError;
} SearchResults;

// load index from filepath
FaissIndex* loadIndex(const char*);

// get metadata using index
FaissMetadata* getMetadata(const FaissIndex*);

SearchResults searchFaiss(const FaissIndex* index, int numVectors, const float* vectors, int topK);

#endif  // __FAISS_WRAPPER_H__
