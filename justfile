# Mere Build System

# Version configuration
VERSION := env_var_or_default("VERSION", "dev")
LDFLAGS := "-ldflags -X main.version=" + VERSION

# Build configuration
BINARY := "mere"
MAIN := "./cmd/mere"
BUILD_DIR := "./bin"

# Default target
default: build

# Display this help message
help:
    @just --list

# Build the mere binary to ./bin/
build:
    @echo "Building {{BINARY}}..."
    mkdir -p "{{BUILD_DIR}}"
    go build {{LDFLAGS}} -o "{{BUILD_DIR}}/{{BINARY}}" {{MAIN}}
    @echo "Built {{BUILD_DIR}}/{{BINARY}}"

# Build binary in current directory
build-local:
    @echo "Building {{BINARY}}..."
    go build "-ldflags=-X main.version={{VERSION}}" -o "{{BINARY}}" {{MAIN}}
    @echo "Built {{BINARY}}"

# Run tests with coverage
test:
    @echo "Running tests..."
    go test -v -race -coverprofile=coverage.out ./...
    @echo "Test coverage saved to coverage.out"

# Run tests and show coverage report
test-coverage: test
    @echo ""
    @echo "Coverage Report:"
    go tool cover -func=coverage.out

# Install to GOBIN
install:
    @echo "Installing {{BINARY}}..."
    go build {{LDFLAGS}} -o "$GOPATH/bin/{{BINARY}}" {{MAIN}}
    @echo "Installed to $GOPATH/bin/{{BINARY}}"

# Download and tidy dependencies
deps:
    @echo "Downloading dependencies..."
    go mod download
    go mod tidy

# Clean build artifacts
clean:
    @echo "Cleaning..."
    rm -rf "{{BUILD_DIR}}"
    rm -f "{{BINARY}}"
    rm -f coverage.out
    @echo "Cleaned"

# Run linter (requires golangci-lint)
lint:
    @echo "Linting..."
    golangci-lint run ./...

# Format code
fmt:
    @echo "Formatting..."
    go fmt ./...

# Run development workflow (clean + deps + build + test)
dev: clean deps build test
    @echo "Development workflow complete"

# Release build with version
release version:
    VERSION="{{version}}" just build-local
