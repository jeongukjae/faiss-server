load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/jeongukjae/faiss-server
gazelle(
    name = "gazelle",
    args = [
        "-go_grpc_compiler=@io_bazel_rules_go//proto:go_grpc",
        "-go_grpc_compiler=@com_github_grpc_ecosystem_grpc_gateway_v2//protoc-gen-grpc-gateway:go_gen_grpc_gateway",
    ],
    command = "update",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_binary(
    name = "faiss-server",
    embed = [":faiss-server_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "faiss-server_lib",
    srcs = ["main.go"],
    importpath = "github.com/jeongukjae/faiss-server",
    visibility = ["//visibility:private"],
    deps = [
        "//protos/faiss_server:service",
        "@com_github_golang_glog//:glog",
        "@com_github_grpc_ecosystem_grpc_gateway_v2//runtime",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
