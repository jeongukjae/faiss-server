load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@com_github_grpc_ecosystem_grpc_gateway_v2//protoc-gen-openapiv2:defs.bzl", "protoc_gen_openapiv2")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "service_proto",
    srcs = ["service.proto"],
    deps = [
        "@go_googleapis//google/api:annotations_proto",
    ],
)

go_proto_library(
    name = "service_go_proto",
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
        "@com_github_grpc_ecosystem_grpc_gateway_v2//protoc-gen-grpc-gateway:go_gen_grpc_gateway",
    ],
    importpath = "github.com/jeongukjae/faiss-server/protos/faiss/service",
    proto = ":service_proto",
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
    ],
)

go_library(
    name = "service",
    embed = [":service_go_proto"],
    importpath = "github.com/jeongukjae/faiss-server/protos/faiss/service",
)

protoc_gen_openapiv2(
    name = "openapi",
    proto = ":service_proto",
)
