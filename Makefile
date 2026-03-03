.PHONY: help fmt test vet lint build

GO ?= go
CMD_PATH ?= ./cmd/shici
BINARY ?= shici
GOCACHE ?= $(CURDIR)/.cache/go-build

help:
	@echo "Available targets:"
	@echo "  make fmt    - format Go code"
	@echo "  make test   - run unit tests"
	@echo "  make vet    - run go vet"
	@echo "  make lint   - run golangci-lint"
	@echo "  make build  - build binary to ./bin/shici"

fmt:
	GOCACHE=$(GOCACHE) $(GO) fmt ./...

test:
	GOCACHE=$(GOCACHE) $(GO) test ./...

vet:
	GOCACHE=$(GOCACHE) $(GO) vet ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed"; exit 127; }
	GOCACHE=$(GOCACHE) golangci-lint run

build:
	mkdir -p bin
	GOCACHE=$(GOCACHE) $(GO) build -o bin/$(BINARY) $(CMD_PATH)
