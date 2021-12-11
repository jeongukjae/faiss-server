#ifndef __FAISS_WRAPPER_H__
#define __FAISS_WRAPPER_H__

#include <stdlib.h>
#include "faiss/c_api/index_c.h"
#include "faiss/c_api/index_io_c.h"

FaissIndex* loadIndex(const char* path);

#endif  // __FAISS_WRAPPER_H__
