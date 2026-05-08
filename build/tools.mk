#
# build/tools.mk - Containerized tool execution (no host install required)
#
# Tools run inside the shared go-ci-tools Docker image via TOOLS_CMD.
# Nothing is installed on the user's system unless USE_LOCAL_TOOLS=1.
#
# Prerequisite: the shared image must be available at TOOLS_IMAGE_FULL.
#   docker pull registry.znet/znet/go-ci-tools:latest
#
# Usage in other fragments: $(RUN_TOOL) <tool> [args...]
# Set USE_LOCAL_TOOLS=1 to run tools installed on the host instead.
#

GO           ?= go
VENDOR_CMD   ?= $(GO) mod tidy
TOOL_DIR     ?= tools

TOOLS_IMAGE      ?= znet/go-ci-tools
TOOLS_IMAGE_TAG  ?= latest
TOOLS_MOUNT_PATH ?= /workspace

TOOLS_IMAGE_FULL = $(if $(registry),$(registry)/$(TOOLS_IMAGE),$(TOOLS_IMAGE))
TOOLS_CMD = docker run --rm -t -v $(abspath .):$(TOOLS_MOUNT_PATH) -w $(TOOLS_MOUNT_PATH) $(TOOLS_IMAGE_FULL):$(TOOLS_IMAGE_TAG)

RUN_TOOL := $(TOOLS_CMD)
ifeq ($(USE_LOCAL_TOOLS),1)
RUN_TOOL :=
endif

# Install tools on the host (only needed when USE_LOCAL_TOOLS=1).
.PHONY: tools tools-tidy
tools:
	@echo "=== $(PROJECT_NAME) === [ tools            ]: installing tools on host..."
	@cd $(TOOL_DIR) && $(GO) install $$(go list -e -tags tools -f '{{join .Imports " "}}' tools.go)
	@cd $(TOOL_DIR) && $(VENDOR_CMD)

tools-tidy:
	@cd $(TOOL_DIR) && $(VENDOR_CMD)
