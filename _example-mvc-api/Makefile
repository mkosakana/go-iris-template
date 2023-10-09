.PHONY: setup
setup:
	@make build
	@make up

.PHONY: build
build:
	docker-compose build --no-cache --force-rm app

.PHONY: up
up:
	@make start_local_db
	docker-compose up app

.PHONY: start_local_db
start_local_db:
	docker compose up -d mysql

.PHONY: down
down:
	docker-compose down

.PHONY: ps
ps:
	docker-compose ps

.PHONY: generate
generate:
	find ./service/ -name "mock_*.go" | xargs rm -f
	go install github.com/golang/mock/mockgen@v1.6.0
	go install golang.org/x/tools/cmd/goimports@latest
	go generate ./... > /dev/null
