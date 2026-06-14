.PHONY: run
run:
	go run ./cmd/api

.PHONY: build
build:
	go build -o bin/api ./cmd/api

.PHONY: dev
dev:
	air  

.PHONY: test
test:
	go test ./...