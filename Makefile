.PHONY: help build test test-cover test-race clean fmt vet lint mod-tidy

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the project
	go build -v ./...

test: ## Run tests
	go test -v ./...

test-cover: ## Run tests with coverage
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-race: ## Run tests with race detection
	go test -v -race ./...

clean: ## Clean build artifacts
	go clean ./...
	rm -f coverage.out coverage.html

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: ## Run golangci-lint (requires golangci-lint to be installed)
	golangci-lint run

mod-tidy: ## Tidy go modules
	go mod tidy

check: test vet ## Run all checks (tests + vet)

ci: test-cover vet ## Run CI checks (coverage + vet)
