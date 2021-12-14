load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

cc_library(
    name = "faiss_wrapper",
    srcs = [
        "wrapper.cpp",
        "wrapper.h",
    ],
    linkopts = [
        "-lblas",
        "-llapack",
        "-lomp",
    ],
    linkstatic = 1,
    deps = [
        "@com_github_facebookresearch_faiss//:faiss_c",
        "@com_github_googleapis_google_cloud_cpp//:storage",
    ],
)

cc_test(
    name = "faiss_wrapper_test",
    srcs = ["wrapper_test.cpp"],
    deps = [
        ":faiss_wrapper",
        "@com_google_googletest//:gtest_main",
    ],
)

go_library(
    name = "faiss",
    srcs = ["faiss.go"],  # keep
    cdeps = [
        ":faiss_wrapper",
    ],
    cgo = True,
    importpath = "github.com/jeongukjae/faiss-server/faiss",
)

go_test(
    name = "faiss_test",
    srcs = ["faiss_test.go"],
    data = ["//testdata:random_index"],
    embed = [":faiss"],
    deps = ["@com_github_stretchr_testify//assert"],
)
