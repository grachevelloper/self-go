.PHONY: run
run:
	go run ./cmd/api

.PHONY: build
build:
	go build -o bin/api ./cmd/api

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go generate ./...
