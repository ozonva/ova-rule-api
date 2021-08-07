.PHONY: build, run, lint, test

build:
	go build -o ./bin/app ./cmd/ova-rule-api

run:
	go run ./cmd/ova-rule-api

lint:
	golangci-lint run -v

test:
	go test ./...
