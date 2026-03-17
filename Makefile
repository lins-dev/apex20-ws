.PHONY: run build test lint setup

run:
	go run ./cmd/ws/...

build:
	go build -o bin/ws ./cmd/ws/...

test:
	go test ./...

lint:
	golangci-lint run ./...

setup:
	git config core.hooksPath .githooks
