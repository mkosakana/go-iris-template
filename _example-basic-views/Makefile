.PHONY: setup
setup:
	@make build
	@make up

.PHONY: build
build:
	docker-compose build --no-cache --force-rm app

.PHONY: up
up:
	docker-compose up app

.PHONY: down
down:
	docker-compose down

.PHONY: ps
ps:
	docker-compose ps
