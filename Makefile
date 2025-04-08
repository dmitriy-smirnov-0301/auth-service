.PHONY: build test lint fmt run

build:
	go build -v ./...

test:
	go test -v ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

run-auth-server:
	go run ./cmd/auth-server/main.go
