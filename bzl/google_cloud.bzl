load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _google_cloud_deps_impl(ctx): # Corrected implementation function name

    # build_file_content = """go_library(name = "go_default_library", srcs = glob(["**/*.go"]), importpath = "cloud.google.com/go/secretmanager", visibility = ["//visibility:public"])"""

    http_archive(
        name = "com_google_cloud_go_secretmanager",
        sha256 = "5e9d2fd3332aa58e20efa4440453d9d5130a5b535b95ce8d3d3b0d10883dcb5e",
        urls = ["file:///Users/parth/Documents/backend-api/patched_secretmanager.zip"],
        build_file_content = """
load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = glob(["**/*.go"]),
    importpath = "cloud.google.com/go/secretmanager",
    visibility = ["//visibility:public"],
)
""",
    )
    http_archive(
        name = "com_google_cloud_go_spanner",
        urls = ["https://proxy.golang.org/cloud.google.com/go/spanner/@v/v1.73.0.zip"],
        sha256 = "2acf166842121d138df06bbbe3a3da15f9b77be57202e9936d21f3897354411d", # Correct SHA
        build_file_content = """
load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = glob(["**/*.go"]),
    importpath = "cloud.google.com/go/spanner",
    visibility = ["//visibility:public"],
)
""",
    )

google_cloud_deps = module_extension(
    implementation = _google_cloud_deps_impl, # Correct reference to the implementation function
)