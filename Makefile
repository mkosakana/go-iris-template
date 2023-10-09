.PHONY: docker-clean
docker-clean:
	docker system prune --volumes -fa && docker builder prune -f

.PHONY: setup-views
setup-views:
	cd _example-basic-views/ && make setup

.PHONY: setup-api
setup-api:
	cd _example-mvc-api/ && make setup
