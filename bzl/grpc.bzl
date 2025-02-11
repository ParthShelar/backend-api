load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _grpc_deps_impl(ctx):
    http_archive(
        name = "org_golang_google_grpc",
        urls = ["https://proxy.golang.org/google.golang.org/grpc/@v/v1.69.2.zip"], # Correct URL
        sha256 = "02fe82fe5b0ba6a6a377fd75776685270367fb66bc357c7185eafd342b212d9a", # Correct SHA256
        build_file_content = """
load("@rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = glob(["**/*.go"]),
    importpath = "google.golang.org/grpc",
    visibility = ["//visibility:public"],
)
""",
    )

grpc_deps = module_extension(
    implementation = _grpc_deps_impl,
)