load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//visibility:public"])

cc_library(
    name = "faiss_wrapper",
    srcs = [
        "faiss_wrapper.c",
        "faiss_wrapper.h",
    ],
    linkopts = [
        "-lblas",
        "-llapack",
        "-lomp",
    ],
    deps = [
        "@com_github_facebookresearch_faiss//:faiss_c",
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
