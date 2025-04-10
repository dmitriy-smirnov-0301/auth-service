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

SERVICE ?= user
VERSION ?= v1

PROTO_DIR = proto
OUT_DIR = pkg/proto

PROTO_PATH = $(PROTO_DIR)/$(SERVICE)/$(VERSION)
GEN_PATH = $(OUT_DIR)/$(SERVICE)/$(VERSION)

generate:
	mkdir -p $(GEN_PATH)
	protoc -I=$(PROTO_PATH) \
		--go_out=$(GEN_PATH) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_PATH) --go-grpc_opt=paths=source_relative \
		$(PROTO_PATH)/*.proto

generate-user:
	make generate SERVICE=user VERSION=v1
