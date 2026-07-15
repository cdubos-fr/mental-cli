version := `git describe --tags --always --dirty 2>/dev/null || echo dev`
ldflags := "-X github.com/cdubos-fr/mental-cli/internal/cli.version=" + version

# List available recipes
default:
    @just --list

# Build the mental binary into ./bin
build:
    mkdir -p bin
    go build -ldflags '{{ ldflags }}' -o bin/mental ./cmd/mental

# Build and run mental, forwarding args: just run dump "message"
run *args:
    go run -ldflags '{{ ldflags }}' ./cmd/mental {{ args }}

# Run the test suite
test:
    go test ./...

# Run the test suite with coverage report
cover:
    go test -covermode=atomic -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

# Format code (gofmt + goimports, as configured in .golangci.yml)
fmt:
    golangci-lint fmt

# Check formatting without writing changes
fmt-check:
    golangci-lint fmt -d

# Vet the code
vet:
    go vet ./...

# Run the linter
lint:
    golangci-lint run ./...

# Tidy go.mod/go.sum
tidy:
    go mod tidy

# Lint GitHub Actions workflows for security issues
security:
    zizmor .github/workflows/

# Run fmt-check, vet, lint, test and security — what CI runs
check: fmt-check vet lint test security

# Build a local snapshot release with GoReleaser (no publish)
release-snapshot:
    goreleaser release --snapshot --clean

# Remove build artifacts
clean:
    rm -rf bin dist coverage.out
