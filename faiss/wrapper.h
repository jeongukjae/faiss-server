#ifndef __FAISS_WRAPPER_H__
#define __FAISS_WRAPPER_H__

#ifdef __cplusplus
#include <string>

extern "C" {
#endif

#include <stdlib.h>

#include "faiss/c_api/Index_c.h"
#include "faiss/c_api/index_io_c.h"

typedef struct SearchResults {
  int64_t* ids;
  float* distances;
  int isError;
} SearchResults;

const char* getError();

// load index from filepath
FaissIndex* loadIndex(const char*);

SearchResults searchFaiss(const FaissIndex* index, int numVectors, int topK,
                          const float* vectors);

int removeVectors(FaissIndex* index, int numIds, const int64_t* ids);

#ifdef __cplusplus
}

// for testing purpose
// below functions are not exposed for cgo binding.

void parseUrl(const std::string url, std::string& bucket,
              std::string& blobPath);
#endif

#endif  // __FAISS_WRAPPER_H__
