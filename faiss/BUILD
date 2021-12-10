load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "faiss",
    srcs = ["faiss.go"],
    cdeps = [
        "@com_github_facebookresearch_faiss//:faiss_c",
    ],
    cgo = True,
    importpath = "github.com/jeongukjae/faiss-server/faiss",
    visibility = ["//visibility:public"],
)
