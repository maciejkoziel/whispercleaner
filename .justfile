# Justfile for Whispercleaner

# Default recipe to run when just is called without arguments
default:
    @just --list

# Build the Whispercleaner executable for the current platform
build:
    go build -o whispercleaner

# Run tests
test:
    go test ./...

# Build for Linux
build-linux:
    GOOS=linux GOARCH=amd64 go build -o whispercleaner-linux-amd64

# Build for Windows
build-windows:
    GOOS=windows GOARCH=amd64 go build -o whispercleaner-windows-amd64.exe

# Build for macOS
build-macos:
    GOOS=darwin GOARCH=amd64 go build -o whispercleaner-darwin-amd64

# Build for all platforms
build-all: build-linux build-windows build-macos

# Clean up built executables
clean:
    rm -f whispercleaner whispercleaner-linux-amd64 whispercleaner-windows-amd64.exe whispercleaner-darwin-amd64

# Run the application (assumes it's built)
run *ARGS:
    ./whispercleaner {{ARGS}}

# Build and run the application
build-and-run *ARGS: build
    ./whispercleaner {{ARGS}}

# Version of the application
version:
    @go run whispercleaner.go --version