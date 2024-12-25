.PHONY: build run test docker

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server -env dev

test:
	go test -v ./...

docker:
	docker build -t dingtalk-service .

docker-run:
	docker run -p 8080:8080 --env-file .env dingtalk-service 