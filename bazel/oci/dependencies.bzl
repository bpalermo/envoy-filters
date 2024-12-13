load("@rules_oci//oci:pull.bzl", "oci_pull")

def oci_dependencies():
    oci_pull(
        name = "distroless_base",
        digest = "sha256:6d4a4f40e93615df1677463ca56456379cc3a4e2359308c9e72bc60ffc4a12a9",
        image = "gcr.io/distroless/base",
        platforms = [
            "linux/amd64",
            "linux/arm64/v8",
        ],
    )
