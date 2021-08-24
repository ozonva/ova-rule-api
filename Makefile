.PHONY: build, run, lint, test, mocks

build:
	go build -o ./bin/app ./cmd/ova-rule-api

run:
	go run ./cmd/ova-rule-api

lint:
	golangci-lint run -v

test: mocks
	go test -race ./...

mocks:
	rm -rf ./internal/mocks/mock_*
	mockgen -source=./internal/repo/repo.go -destination=./internal/mocks/mock_repo.go -package mocks
	mockgen -source=./internal/flusher/flusher.go -destination=./internal/mocks/mock_flusher.go -package mocks
