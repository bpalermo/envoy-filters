load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# 1. Determine SHA256 `wget https://github.com/envoyproxy/envoy/archive/$COMMIT.tar.gz && sha256sum $COMMIT.tar.gz`
# 2. Update .bazelversion, envoy.bazelrc and .bazelrc if needed.
#
# Commit date: 2024-12-08
ENVOY_COMMIT = "a0504e87c5a246cb097b37049b1e4dc7706c2a90"

ENVOY_SHA256 = "60c34f0d91caf66a87175a091dcee2c00715001bdd5e6f6f80079179e0abc636"

ENVOY_ORG = "envoyproxy"

ENVOY_REPO = "envoy"

def envoy_repositories():
    http_archive(
        name = "envoy",
        sha256 = ENVOY_SHA256,
        strip_prefix = ENVOY_REPO + "-" + ENVOY_COMMIT,
        url = "https://github.com/" + ENVOY_ORG + "/" + ENVOY_REPO + "/archive/" + ENVOY_COMMIT + ".tar.gz",
    )
