/*
This file is copied from below link and slightly modified to dump random index.

https://github.com/facebookresearch/faiss/blob/d68ff421957a02f6137796cb53e292ed24f5a369/tutorial/cpp/1-Flat.cpp
*/

#include <unistd.h>

#include <cstdio>
#include <cstdlib>
#include <random>
#include <string>

#include "faiss/Index.h"
#include "faiss/IndexFlat.h"
#include "faiss/index_io.h"

// 64-bit int
using idx_t = faiss::Index::idx_t;

#define INDEX_OUTPUT_FILENAME "testdata/random-index.faiss"
#define RANDOM_SEED 1234

int main() {
  int d = 8;     // dimension
  int nb = 1000;  // database size
  int nq = 1000;  // nb of queries

  std::mt19937 rng;
  rng.seed(RANDOM_SEED);
  std::uniform_real_distribution<> distrib;

  float* xb = new float[d * nb];
  float* xq = new float[d * nq];

  for (int i = 0; i < nb; i++) {
    for (int j = 0; j < d; j++) xb[d * i + j] = distrib(rng);
    xb[d * i] += i / 1000.;
  }

  for (int i = 0; i < nq; i++) {
    for (int j = 0; j < d; j++) xq[d * i + j] = distrib(rng);
    xq[d * i] += i / 1000.;
  }

  faiss::IndexFlatL2 index(d);  // call constructor
  printf("is_trained = %s\n", index.is_trained ? "true" : "false");
  index.add(nb, xb);  // add vectors to the index
  printf("ntotal = %lld\n", index.ntotal);
  faiss::write_index(reinterpret_cast<const faiss::Index*>(&index),
                     INDEX_OUTPUT_FILENAME);

  {
    char buff[FILENAME_MAX];
    getcwd(buff, FILENAME_MAX);
    printf("Random Index path: %s/%s\n\n", buff, INDEX_OUTPUT_FILENAME);
  }

  int k = 4;

  {  // sanity check: search 5 first vectors of xb
    idx_t* I = new idx_t[k * 5];
    float* D = new float[k * 5];

    index.search(5, xb, k, D, I);
    printf("query vector=\n");
    for (int i = 0; i < 5; i++) {
      for (int j = 0; j < d; j++) {
        printf(" %5g", xb[i * d + j]);
      }
      printf("\n");
    }

    // print results
    printf("I=\n");
    for (int i = 0; i < 5; i++) {
      for (int j = 0; j < k; j++) printf("%5lld ", I[i * k + j]);
      printf("\n");
    }

    printf("D=\n");
    for (int i = 0; i < 5; i++) {
      for (int j = 0; j < k; j++) printf("%7g ", D[i * k + j]);
      printf("\n");
    }

    delete[] I;
    delete[] D;
  }

  {  // search xq
    idx_t* I = new idx_t[k * nq];
    float* D = new float[k * nq];

    index.search(nq, xq, k, D, I);

    // print results
    printf("I (5 first results)=\n");
    for (int i = 0; i < 5; i++) {
      for (int j = 0; j < k; j++) printf("%5lld ", I[i * k + j]);
      printf("\n");
    }

    printf("I (5 last results)=\n");
    for (int i = nq - 5; i < nq; i++) {
      for (int j = 0; j < k; j++) printf("%5lld ", I[i * k + j]);
      printf("\n");
    }

    delete[] I;
    delete[] D;
  }

  delete[] xb;
  delete[] xq;

  return 0;
}
