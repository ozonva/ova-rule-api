export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

all: generate test build

.PHONY: deps
deps:
	go get -u github.com/pkg/errors
	go get -u github.com/stretchr/testify
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u gopkg.in/yaml.v2
	go get -u github.com/golang/mock
	go get -u github.com/rs/zerolog/log
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/protobuf/reflect/protoreflect
	go get -u google.golang.org/protobuf/runtime/protoimpl
	go get -u github.com/jackc/pgx/v4
	go get -u github.com/georgysavva/scany
	go get -u github.com/Masterminds/squirrel
	go get -u github.com/Shopify/sarama
	go get -u github.com/uber/jaeger-client-go
	go get -u github.com/prometheus/client_golang
	go get -u github.com/spf13/viper
	go get -u golang.org/x/tools/cmd/godoc

	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: generate
generate:
	protoc --proto_path=. -I vendor.protogen \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	api/api.proto

.PHONY: build
build: deps
	go build -o ./bin/app ./cmd/ova-rule-api

.PHONY: run
run:
	go run ./cmd/ova-rule-api/main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run ./... --fix

.PHONY: test
test: mocks
	go test -race ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

.PHONY: mocks
mocks:
	rm -rf ./internal/mocks/mock_*
	mockgen -source=./internal/repo/repo.go -destination=./internal/mocks/mock_repo.go -package mocks
	mockgen -source=./internal/flusher/flusher.go -destination=./internal/mocks/mock_flusher.go -package mocks
	mockgen -source=./internal/kafka/producer.go -destination=./internal/mocks/mock_producer.go -package mocks
	mockgen -source=./internal/metrics/metrics.go -destination=./internal/mocks/mock_metrics.go -package mocks

.PHONY: migrate-up
migrate-up:
	cd migrations &&\
 	goose postgres "postgres://ova:iloveozon@localhost:5432/ova?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	cd migrations &&\
 	goose postgres "postgres://ova:iloveozon@localhost:5432/ova?sslmode=disable" down

.PHONY: migrate-status
migrate-status:
	cd migrations &&\
 	goose postgres "postgres://ova:iloveozon@localhost:5432/ova?sslmode=disable" status

.PHONY: up
up:
	docker-compose up

.PHONY: down
down:
	docker-compose down

.PHONY: doc
doc:
	godoc -http=:6060
