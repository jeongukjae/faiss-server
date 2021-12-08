load("@rules_foreign_cc//foreign_cc:defs.bzl", "cmake")

package(default_visibility = ["//visibility:public"])

licenses(["notice"])  # MIT

exports_files(["LICENSE"])

filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
)

cmake(
    name = "faiss_c_api",
    lib_source = ":all_srcs",
    generate_args = [
        "-G Ninja",
        "-DFAISS_ENABLE_GPU=OFF",
        "-DFAISS_ENABLE_PYTHON=OFF",
        "-DBUILD_TESTING=OFF",
        "-DBUILD_SHARED_LIBS=OFF",
        "-DFAISS_ENABLE_C_API=ON",
        "-DCMAKE_BUILD_TYPE=Release",
        "-DFAISS_OPT_LEVEL=avx2",
    ],
    out_static_libs = ["libfaiss.a"],
)
