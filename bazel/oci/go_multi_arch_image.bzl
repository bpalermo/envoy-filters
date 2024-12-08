load("@aspect_bazel_lib//lib:testing.bzl", "assert_archive_contains")
load("@container_structure_test//:defs.bzl", "container_structure_test")
load("@rules_pkg//:pkg.bzl", "pkg_tar")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_image_index", "oci_load", "oci_push")
load("//bazel/oci:transition.bzl", "multi_arch")

def go_multi_arch_image(name, repository, base = "@distroless_base", container_test_configs = ["testdata/container_test.yaml"], tars = [], **kwargs):
    """
    Creates a containerized binary from Go sources.
    Parameters:
        name:  name of the image
        repository: image repository
        base: base image
        tars: additional image layers
        kwargs: arguments passed to the go_binary target
    """

    pkg_tar(
        name = "layer",
        srcs = [name],
        package_dir = "/",
        visibility = ["//visibility:private"],
    )

    assert_archive_contains(
        name = "test_layer",
        archive = "layer.tar",
        expected = [name],
        visibility = ["//visibility:private"],
    )

    oci_image(
        name = "image",
        base = "@distroless_base",
        entrypoint = ["/{}".format(name)],
        tars = [":layer"],
        visibility = ["//visibility:private"],
    )

    container_structure_test(
        name = "test_image",
        configs = container_test_configs,
        image = ":image",
        tags = ["requires-docker"],
        visibility = ["//visibility:private"],
    )

    multi_arch(
        name = "images",
        image = ":image",
        platforms = [
            "@io_bazel_rules_go//go/toolchain:linux_amd64",
            "@io_bazel_rules_go//go/toolchain:linux_arm64",
        ],
        visibility = ["//visibility:private"],
    )

    oci_image_index(
        name = "index",
        images = [
            ":images",
        ],
    )

    oci_load(
        name = "load",
        image = ":index",
        repo_tags = ["latest"],
        format = "oci",
    )

    oci_push(
        name = "push",
        image = ":index",
        repository = repository,
    )