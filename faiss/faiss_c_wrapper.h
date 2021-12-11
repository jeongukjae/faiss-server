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

// load index from filepath
FaissIndex* loadIndex(const char*);

// get metadata using index
FaissMetadata* getMetadata(FaissIndex*);

#endif  // __FAISS_WRAPPER_H__
