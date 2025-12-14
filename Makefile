RUN_CMD = go run ./cmd/testflowkit/*.go run
APP_NAME := tkit
SRC_DIR := ./cmd/testflowkit/*.go
BUILD_DIR := ./build
VERSION := $(shell git describe --tags --always --dirty)
# Only include supported architecture combinations
GOARCHS := amd64 arm64
GOOSES := linux darwin windows
LDFLAGS := -X 'main.Version=$(VERSION)'
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
NUM_CPUS := $(shell nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 4)

# Go build flags for optimization
GOBUILD := go build -trimpath -ldflags "$(LDFLAGS) -X 'main.BuildTime=$(BUILD_TIME)'"

.PHONY: all
all: lint test build

.PHONY: lint
lint:
	@echo "Running linter..."
	@golangci-lint run

.PHONY: run
run:
	${RUN_CMD}

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | grep total:

.PHONY: run_e2e
run_e2e:
	${RUN_CMD} -c e2e/config.yml

.PHONY: run_e2e_server
run_e2e_server:
	go run e2e/server/main.go

.PHONY: generate_doc
generate_doc:
	go run scripts/doc_generator/main.go

.PHONY: clean
clean:
	@echo "Cleaning up build directory..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out
	@echo "Done."

# Build for all OS/Arch combinations
.PHONY: releases
releases:
	@echo "Building for all known architectures..."
	@mkdir -p $(BUILD_DIR)
	@start_time=$$(date +%s); \
	success_count=0; \
	fail_count=0; \
	jobs=""; \
	for os in $(GOOSES); do \
		for arch in $(GOARCHS); do \
			echo "Building for $$os/$$arch..."; \
			$(MAKE) build GOOS=$$os GOARCH=$$arch & \
			jobs="$$jobs $$!"; \
		done \
	done; \
	for job in $$jobs; do \
		if wait $$job; then \
			success_count=$$(expr $$success_count + 1); \
		else \
			fail_count=$$(expr $$fail_count + 1); \
		fi \
	done; \
	end_time=$$(date +%s); \
	duration=$$(expr $$end_time - $$start_time); \
	echo "Build Summary:"; \
	echo " 	Successful builds: $$success_count"; \
	echo " 	Failed builds: $$fail_count"; \
	echo " 	Total time: $$duration seconds"; \
	test $$fail_count -eq 0

# Build for a specific OS/Arch
.PHONY: build
build:
	@if [ -z "$(GOOS)" ]; then \
		echo "Error: GOOS is not set"; \
		exit 1; \
	fi
	@if [ -z "$(GOARCH)" ]; then \
		echo "Error: GOARCH is not set"; \
		exit 1; \
	fi
	@echo "Building for GOOS=$(GOOS) GOARCH=$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	@start_time=$$(date +%s); \
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME)$(if $(findstring windows,$(GOOS)),.exe,) $(SRC_DIR) || { \
		echo "Build failed for $(GOOS)/$(GOARCH)"; \
		exit 1; \
	}; \
	cd $(BUILD_DIR) && zip -q $(APP_NAME)-$(GOOS)-$(GOARCH).zip $(APP_NAME)$(if $(findstring windows,$(GOOS)),.exe,) && cd ..; \
	rm -f $(BUILD_DIR)/$(APP_NAME)$(if $(findstring windows,$(GOOS)),.exe,); \
	end_time=$$(date +%s); \
	duration=$$((end_time - start_time)); \
	echo "Build complete: $(BUILD_DIR)/$(APP_NAME)-$(GOOS)-$(GOARCH).zip ($$duration seconds)"

.PHONY: help
help:
	@echo "Makefile for building Go applications"
	@echo
	@echo "Targets:"
	@echo "  all               Run lint, test, and build"
	@echo "  lint             Run golangci-lint"
	@echo "  run              Run testflowkit app"
	@echo "  test             Run tests with coverage"
	@echo "  releases         Build for all architectures in parallel"
	@echo "  build            Build for a specific OS and architecture"
	@echo "                  Example: make build GOOS=linux GOARCH=amd64"
	@echo "  clean            Remove all build artifacts"
	@echo "  help             Show this help message"
	@echo
	@echo "Build Information:"
	@echo "  Version: $(VERSION)"
	@echo "  Build Time: $(BUILD_TIME)"
	@echo "  Parallel Builds: $(NUM_CPUS)"

