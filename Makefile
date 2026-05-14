.PHONY: build test run up down

build:
	go build -o bin/veilid ./cmd/veilid/main.go

test:
	go test -v ./internal/...

run:
	go run cmd/veilid/main.go

up:
	docker-compose up -d

down:
	docker-compose down