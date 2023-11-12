.PHONY: docker-clean
docker-clean:
	docker system prune --volumes -fa && docker builder prune -f

.PHONY: setup-views
setup-views:
	cd _example-basic-views/ && make setup

.PHONY: setup-api
setup-api:
	cd _example-mvc-api/ && make setup

.PHONY: test_for_ci
test_for_ci:
	cd _example-basic-views/ &&\
	go test ./... &&\
	cd ../_example-mvc-api/ &&\
	go test ./...

.PHONY: fmt_for_ci
fmt_for_ci:
	cd _example-basic-views/ &&\
    	go fmt ./... &&\
    	cd ../_example-mvc-api/ &&\
    	go fmt ./...

.PHONY: lint_for_ci
lint_for_ci:
	cd _example-basic-views/ &&\
    	go vet ./... &&\
    	cd ../_example-mvc-api/ &&\
    	go vet ./...
