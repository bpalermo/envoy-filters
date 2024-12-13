load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

RULES_OCI_VERSION = "2.0.1"
RULES_OCI_SHA256 = "acbf8f40e062f707f8754e914dcb0013803c6e5e3679d3e05b571a9f5c7e0b43"

CONTAINER_STRUCTURE_TEST_VERSION = "1.19.3"
CONTAINER_STRUCTURE_TEST_SHA256 = "c91a76f7b4949775941f8308ee7676285555ae4756ec1ec990c609c975a55f93"

def oci_repositories():
    http_archive(
        name = "rules_oci",
        sha256 = RULES_OCI_SHA256,
        strip_prefix = "rules_oci-{}".format(RULES_OCI_VERSION),
        url = "https://github.com/bazel-contrib/rules_oci/releases/download/v{}/rules_oci-v{}.tar.gz".format(RULES_OCI_VERSION, RULES_OCI_VERSION),
    )

    http_archive(
        name = "container_structure_test",
        sha256 = CONTAINER_STRUCTURE_TEST_SHA256,
        strip_prefix = "container-structure-test-{}".format(CONTAINER_STRUCTURE_TEST_VERSION),
        url = "https://github.com/GoogleContainerTools/container-structure-test/archive/refs/tags/v{}.tar.gz".format(CONTAINER_STRUCTURE_TEST_VERSION),
    )
