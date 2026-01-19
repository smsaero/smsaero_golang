# Variables
BINARY_NAME=smsaero-golang
VERSION=2.0.0
BUILD_DIR=bin
COVERAGE_DIR=coverage

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet
GOLINT=golangci-lint

# Flags
BUILD_FLAGS=-ldflags "-X main.version=$(VERSION)"
TEST_FLAGS=-v -coverprofile=$(COVERAGE_DIR)/coverage.out
LINT_FLAGS=run --timeout=5m

# Phony targets
.PHONY: all build clean deps fmt vet lint go_lint test test-coverage
.PHONY: staticcheck gosec static-analysis
.PHONY: docker-build docker-run demo
.PHONY: update-deps check-deps docs check-all quick-check full-check help

# Main targets
all: clean deps fmt vet lint static-analysis build

# Build
build:
	@echo "Building library..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./demo/

# Clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -rf $(COVERAGE_DIR)

# Dependencies
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Code verification
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Linting
lint:
	@echo "Running linter..."
	@if command -v $(GOLINT) >/dev/null 2>&1; then \
		$(GOLINT) $(LINT_FLAGS) ./...; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest"; \
	fi

# Alias for CI compatibility
go_lint: lint

# Static analysis
staticcheck:
	@echo "Running staticcheck..."
	@if command -v staticcheck >/dev/null 2>&1; then \
		staticcheck ./...; \
	else \
		echo "staticcheck not installed. Install with: brew install staticcheck"; \
		exit 0; \
	fi

# Security analysis
gosec:
	@echo "Running security analysis..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "gosec not installed. Install with: brew install gosec"; \
		exit 0; \
	fi

# Full static analysis
static-analysis: vet staticcheck gosec
	@echo "Static analysis completed successfully"

# Testing
test:
	@echo "Running tests..."
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) $(TEST_FLAGS) ./...

# Coverage testing
test-coverage: test
	@echo "Generating coverage report..."
	$(GOCMD) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "Coverage report saved to $(COVERAGE_DIR)/coverage.html"

# Installation
install: build
	@echo "Installing library..."
	$(GOCMD) install ./...

# Uninstallation
uninstall:
	@echo "Uninstalling library..."
	$(GOCMD) clean -i

# Docker
docker-build:
	@echo "Building Docker image..."
	docker build -t smsaero-golang:$(VERSION) .

docker-run: docker-build
	@echo "Running Docker container..."
	docker run --rm -it smsaero-golang:$(VERSION)

# Demo
demo: build
	@echo "Running demo..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Update dependencies
update-deps:
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

# Check outdated dependencies
check-deps:
	@echo "Checking for outdated dependencies..."
	@if command -v go-mod-outdated >/dev/null 2>&1; then \
		go-mod-outdated -direct -update; \
	else \
		echo "go-mod-outdated not installed. Install with: go install github.com/psampaz/go-mod-outdated@latest"; \
	fi

# Generate documentation
docs:
	@echo "Generating documentation..."
	@if command -v godoc >/dev/null 2>&1; then \
		godoc -http=:6060; \
	else \
		echo "godoc not installed. Install with: go install golang.org/x/tools/cmd/godoc@latest"; \
	fi

# Check all
check-all: fmt vet lint test

# Quick check
quick-check: fmt vet

# Full check
full-check: clean deps check-all test-coverage

# Help
help:
	@echo "Available commands:"
	@echo ""
	@echo "Main:"
	@echo "  all          - Full build (clean, deps, fmt, vet, lint, build)"
	@echo "  build        - Build library"
	@echo "  clean        - Clean temporary files"
	@echo "  deps         - Install dependencies"
	@echo ""
	@echo "Code quality:"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  lint         - Run linter (golangci-lint)"
	@echo ""
	@echo "Static analysis:"
	@echo "  staticcheck  - Static analysis (staticcheck)"
	@echo "  gosec        - Security analysis (gosec)"
	@echo "  static-analysis - Full static analysis (vet, staticcheck, gosec)"
	@echo ""
	@echo "Testing:"
	@echo "  test         - Run tests with coverage"
	@echo "  test-coverage- Generate HTML coverage report"
	@echo ""
	@echo "Docker:"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo ""
	@echo "Utilities:"
	@echo "  demo         - Run demo"
	@echo "  docs         - Generate documentation"
	@echo "  update-deps  - Update dependencies"
	@echo "  check-deps   - Check outdated dependencies"
	@echo ""
	@echo "Checks:"
	@echo "  check-all    - Full check (fmt, vet, lint)"
	@echo "  quick-check  - Quick check (fmt, vet)"
	@echo "  full-check   - Full check (clean, deps, check-all)"
