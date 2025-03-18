# Build
BIN_NAME := windy-judge
VERSION := 1.0.0
PKGS := $(shell go list ./... | grep -v 'examples')

.PHONY: build
build:
	@go build -o $(BIN_NAME) ./main.go

.PHONY: install
install:
	@go install ./...

.PHONY: test
test:
	@go test -v -cover $(PKGS)

.PHONY: release
test-unit: release
	@chmod +x shell/judge.sh $(BIN_NAME)
	@./shell/judge.sh

.PHONY: lint
lint:
	@gofmt -d .
	@golint ./...

.PHONY: clean
clean:
	@rm -f $(BIN_NAME) $(BIN_NAME)-$(VERSION)-linux-amd64 $(BIN_NAME)-$(VERSION)-darwin-amd64
	@go clean

.PHONY: release
release:
	@CGO_ENABLED=0 go build  -buildvcs=false  -ldflags="-s -w"

.PHONY: all
all: lint test build

.DEFAULT_GOAL := build
