load("@rules_foreign_cc//foreign_cc:defs.bzl", "cmake")

package(default_visibility = ["//visibility:public"])

licenses(["notice"])  # MIT

exports_files(["LICENSE"])

filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
)

cmake(
    name = "faiss_c",
    generate_args = [
        "-G Ninja",
        "-DFAISS_ENABLE_GPU=OFF",
        "-DFAISS_ENABLE_PYTHON=OFF",
        "-DFAISS_ENABLE_C_API=ON",
        "-DBUILD_TESTING=OFF",
        "-DCMAKE_BUILD_TYPE=Release",
        "-DFAISS_OPT_LEVEL=general",
    ],
    lib_source = ":all_srcs",
    out_static_libs = [
        "libfaiss_c.a",
        "libfaiss.a",
    ],
    targets = [
        "faiss_c",
        "faiss",
    ],
)
