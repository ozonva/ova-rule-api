export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

all: generate test build

.PHONY: deps
deps:
	go get -u github.com/stretchr/testify
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u gopkg.in/yaml.v2
	go get -u github.com/golang/mock
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: generate
generate:
	protoc -I vendor.protogen \
	--go_out=pkg/rule --go_opt=paths=import \
	--go-grpc_out=pkg/rule --go-grpc_opt=paths=import \
	api/rule/rule.proto

.PHONY: build
build: deps
	go build -o ./bin/app ./cmd/ova-rule-api

.PHONY: run
run:
	go run ./cmd/ova-rule-api

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: test
test: mocks
	go test -race ./...

.PHONY: mocks
mocks:
	rm -rf ./internal/mocks/mock_*
	mockgen -source=./internal/repo/repo.go -destination=./internal/mocks/mock_repo.go -package mocks
	mockgen -source=./internal/flusher/flusher.go -destination=./internal/mocks/mock_flusher.go -package mocks
