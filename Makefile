all: format tidy lint

tidy:
	@go mod tidy

format:
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -s -w {} +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w  -local github.com/RashadAnsari {} +

lint:
	@golangci-lint -c .golangci.yml run ./...

up:
	@docker-compose up -d zookeeper kafka mysql

connect:
	@docker-compose up -d kafka-connect
	@docker logs -f kafka-connect

generate:
	@go run .

ps:
	@docker-compose ps

down:
	@docker-compose down
