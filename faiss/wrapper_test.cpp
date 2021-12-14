#include "wrapper.h"

#include "gmock/gmock.h"
#include "gtest/gtest.h"

TEST(TestWrapper, parseUrl) {
  std::string url = "gs://bucket_name/path/to/blob.index";
  std::string bucket, blobPath;
  parseUrl(url, bucket, blobPath);

  ASSERT_EQ("bucket_name", bucket);
  ASSERT_EQ("path/to/blob.index", blobPath);
}
