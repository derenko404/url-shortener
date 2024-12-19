include ./configs/.env

.PHONY: run dev build docker

docker:
	docker-compose up --build

dev:
	air -c .air.toml

run:
	go run ./cmd/main.go

build:
	rm -rf ./bin
	go build -o ./bin/main ./cmd/main.go