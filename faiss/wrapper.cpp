#include "wrapper.h"

#include <sstream>
#include <string>

#include "faiss/impl/AuxIndexStructures.h"
#include "faiss/impl/io.h"
#include "faiss/index_io.h"
#include "google/cloud/storage/client.h"

namespace gcs = ::google::cloud::storage;

static std::string error;

const char* getError() { return error.c_str(); }

struct GCSIOReader : faiss::IOReader {
  gcs::Client client;
  const std::string bucket;
  const std::string blobPath;
  size_t offset;

  GCSIOReader(gcs::Client client, const std::string bucket,
              const std::string blobPath)
      : client(client), bucket(bucket), blobPath(blobPath), offset(0) {}

  size_t operator()(void* ptr, size_t size, size_t nitems) {
    size_t bufferSize = size * nitems;
    gcs::ObjectReadStream stream = client.ReadObject(
        bucket, blobPath, gcs::ReadRange(offset, offset + bufferSize));
    if (!stream) {
      std::ostringstream ss;
      ss << "Error reading object: " << stream.status();
      error = ss.str();
      return 0;
    }

    size_t nReads = bufferSize;
    stream.read(static_cast<char*>(ptr), bufferSize);
    if (!stream) {
      nReads = stream.gcount();

      std::ostringstream ss;
      ss << "Cannot read successfully from GCS. offset: " << offset
         << ", expected size: " << bufferSize << ", actual: " << nReads;
      error = ss.str();
    }
    offset += nReads;
    stream.Close();

    return nReads / size;
  }
};

FaissIndex* loadIndex(const char* cPath) {
  std::string path(cPath);

  // Handle GCS
  if (path.find("gs://") == 0) {
    gcs::Client client = gcs::Client();
    std::string bucket, blobPath;
    parseUrl(path, bucket, blobPath);

    if (bucket == "" || blobPath == "") {
      error = "Cannot parse GCS url. url: " + path;
      return NULL;
    }

    GCSIOReader reader(client, bucket, blobPath);
    FaissIndex* index = NULL;
    try {
      index = reinterpret_cast<FaissIndex*>(
          faiss::read_index(&reader, FAISS_IO_FLAG_READ_ONLY));
    } catch (std::exception& e) {
      error = e.what();
      return NULL;
    }

    return index;
  }

  // Hadle local fs
  FaissIndex* index = NULL;
  if (faiss_read_index_fname(cPath, FAISS_IO_FLAG_READ_ONLY, &index)) {
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

void parseUrl(const std::string url, std::string& bucket,
              std::string& blobPath) {
  auto pos = url.find("://");
  if (pos == std::string::npos) return;
  auto posOfFirstSlash = url.find("/", pos + 3);
  if (posOfFirstSlash == std::string::npos) return;

  bucket = url.substr(pos + 3, posOfFirstSlash - (pos + 3));
  blobPath = url.substr(posOfFirstSlash + 1);
}
