workspace(name = "faiss-server")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ====================
# rules go and gazelle
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.17.2")

gazelle_dependencies()

# ================
# Rules foreign cc
http_archive(
    name = "rules_foreign_cc",
    sha256 = "1df78c7d7eed2dc21b8b325a2853c31933a81e7b780f9a59a5d078be9008b13a",
    strip_prefix = "rules_foreign_cc-0.7.0",
    url = "https://github.com/bazelbuild/rules_foreign_cc/archive/0.7.0.tar.gz",
)

load("@rules_foreign_cc//foreign_cc:repositories.bzl", "rules_foreign_cc_dependencies")

rules_foreign_cc_dependencies()

# ===========
# rules proto
http_archive(
    name = "rules_proto",
    sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
    strip_prefix = "rules_proto-4.0.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

rules_proto_dependencies()

rules_proto_toolchains()

# ===============
# third party lib
http_archive(
    name = "com_github_facebookresearch_faiss",
    build_file = "//third_party:faiss.BUILD",
    strip_prefix = "faiss-1.7.1",
    sha256 = "7e53e763f4081f9fb329634bb51cecaad674b22f5ee56351d09b0fe21bbc4f72",
    url = "https://github.com/facebookresearch/faiss/archive/v1.7.1.zip",
)

http_archive(
    name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    sha256 = "732f8882dffcceb7c14839ffcf492ae0f5b5dcf2e79003c4ca1b83f29892483e",
    strip_prefix = "grpc-gateway-2.7.1",
    url = "https://github.com/grpc-ecosystem/grpc-gateway/archive/v2.7.1.zip",
)

load("@com_github_grpc_ecosystem_grpc_gateway_v2//:repositories.bzl", grpc_gateway_go_repositories = "go_repositories")

grpc_gateway_go_repositories()
