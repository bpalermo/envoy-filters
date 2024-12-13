load("//bazel/envoy:repositories.bzl", "envoy_repositories")
load("//bazel/oci:repositories.bzl", "oci_repositories")

def envoy_filters_repositories():
    envoy_repositories()
    oci_repositories()
