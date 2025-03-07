RUN_CMD = go run ./cmd/testflowkit/main.go run
APP_NAME := testflowkit
SRC_DIR := ./cmd/testflowkit
BUILD_DIR := ./build
VERSION := $(shell git describe --tags --always --dirty)
GOARCHS := amd64 arm64 386 arm
GOOSES := linux darwin windows
LDFLAGS := -X 'main.Version=$(VERSION)'
lint:
	golangci-lint run

run:
	${RUN_CMD}

unit_test:
	 go test ./...

run_e2e:
	${RUN_CMD} -f e2e/frontend.yml -c e2e/cli.yml

run_e2e_server:
	go run e2e/server/main.go

generate_doc:
	go run scripts/doc_generator/main.go
	
.PHONY: clean
clean:
	@echo "Cleaning up build directory..."
	@rm -rf $(BUILD_DIR)
	@echo "Done."

# Build for all OS/Arch combinations
.PHONY: releases
releases:
	@echo "Building for all known architectures..."
	@mkdir -p $(BUILD_DIR)
	@for os in $(GOOSES); do \
		for arch in $(GOARCHS); do \
			echo "Building for $$os/$$arch..."; \
			GOOS=$$os GOARCH=$$arch go build cmd/testflowkit/main.go -o $(BUILD_DIR)/$(APP_NAME)-$$os-$$arch -ldflags "$(LDFLAGS)" $(SRC_DIR) || echo "Failed to build for $$os/$$arch"; \
		done \
	done
	@echo "All builds complete."

# Build for a specific OS/Arch
.PHONY: build
build:
	ifndef GOOS
		$(error GOOS is not set)
	endif
	ifndef GOARCH
		$(error GOARCH is not set)
	endif
		@echo "Building for GOOS=$(GOOS) GOARCH=$(GOARCH)..."
		@mkdir -p $(BUILD_DIR)
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME)-$(GOOS)-$(GOARCH) -ldflags "$(LDFLAGS)" $(SRC_DIR)
		@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)-$(GOOS)-$(GOARCH)"

# Usage instructions
.PHONY: help
help:
	@echo "Makefile for building Go applications"
	@echo
	@echo "Targets:"
	@echo "  lint               Default lint codebase"
	@echo "  run				run testflowkit app"
	@echo "  releases     		Build for all architectures (GOARCH and GOOS combinations)"
	@echo "  build              Build for a specific OS and architecture (requires GOOS and GOARCH)"
	@echo "                    Example: make build GOOS=linux GOARCH=amd64"
	@echo "  clean              Remove all build artifacts"
	@echo "  help               Show this help message"

