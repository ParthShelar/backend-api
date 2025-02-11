load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _prometheus_deps_impl(ctx):
    http_archive(
        name = "com_github_prometheus_client_golang",
        urls = ["file:///Users/parth/Documents/backend-api/patched_prometheus.zip"], # Correct URL
        sha256 = "d1bdf4c03b371cacb7eb5c2a1dee321fa998e026010043f62e073662dd9800e3", # Correct SHA256
        strip_prefix = "patched_prometheus/github.com/prometheus/client_golang",
    )

prometheus_deps = module_extension(
    implementation = _prometheus_deps_impl,
)