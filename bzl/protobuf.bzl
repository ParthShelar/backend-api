load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _protobuf_deps_impl(ctx):
    http_archive(
        name = "com_google_protobuf",
        urls = ["https://proxy.golang.org/google.golang.org/protobuf/@v/v1.36.1.zip"],  # Correct URL
        sha256 = "9251166c8970737c771910b750f9f4165704433815f92485a14869456245f989",  # Correct SHA256
    )

protobuf_deps = module_extension(
    implementation = _protobuf_deps_impl,
)