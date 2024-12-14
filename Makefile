.PHONY: load-ext-proc
load-ext-proc:
	@bazel run //ext-proc/cmd/server:load
