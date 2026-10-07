BINARY    := leetcode-go
BUILD_DIR := build

.DEFAULT_GOAL := all
.PHONY: all fmt lint critic test build clean

all: fmt lint critic test build

## fmt: format code with gofumpt + goimports (configured in .golangci.yml)
fmt:
	golangci-lint fmt ./...

## lint: run the full golangci-lint suite
lint:
	golangci-lint run ./...

## critic: run only gocritic
critic:
	golangci-lint run --enable-only gocritic ./...

## test: run tests with the race detector
test:
	go test -race ./...

## build: compile the binary into $(BUILD_DIR)/
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY) .

clean:
	rm -rf $(BUILD_DIR)
