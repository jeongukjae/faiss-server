#include "wrapper.h"

#include <faiss/impl/AuxIndexStructures.h>

FaissIndex* loadIndex(const char* path) {
  FaissIndex* index = NULL;
  if (faiss_read_index_fname(path, FAISS_IO_FLAG_READ_ONLY, &index)) {
    return NULL;
  }
  return index;
}

SearchResults searchFaiss(const FaissIndex* index, int numVectors, int topK,
                          const float* vectors) {
  idx_t* ids = (idx_t*)malloc(sizeof(idx_t) * topK * numVectors);
  float* distances = (float*)malloc(sizeof(float) * topK * numVectors);

  int result =
      faiss_Index_search(index, numVectors, vectors, topK, distances, ids);

  SearchResults searchResult = {
      ids,
      distances,
      result,
  };
  return searchResult;
}

int removeVectors(FaissIndex* index, int numIds, const int64_t* ids) {
  faiss::IDSelectorArray selector(numIds, ids);
  size_t nRemoved;
  int code = faiss_Index_remove_ids(
      index, reinterpret_cast<const FaissIDSelector*>(&selector), &nRemoved);
  if (code == 0) return (int)nRemoved;

  // errors
  return -1;
}
