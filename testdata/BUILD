package(default_visibility = ["//visibility:public"])

cc_binary(
    name = "generate_flat_index",
    testonly = True,
    srcs = ["generate_flat_index.cc"],
    linkopts = [
        "-lblas",
        "-llapack",
        "-lomp",
    ],
    deps = ["@com_github_facebookresearch_faiss//:faiss_c"],
)

filegroup(
    name = "random_index",
    srcs = ["random-index.faiss"],
)
