#ifndef __FAISS_WRAPPER_H__
#define __FAISS_WRAPPER_H__

#include <stdlib.h>

#include "faiss/c_api/Index_c.h"
#include "faiss/c_api/index_io_c.h"

typedef struct SearchResults {
  int64_t* ids;
  float* distances;
  int isError;
} SearchResults;

// load index from filepath
FaissIndex* loadIndex(const char*);

SearchResults searchFaiss(const FaissIndex* index, int numVectors, int topK,
                          const float* vectors);

#endif  // __FAISS_WRAPPER_H__
