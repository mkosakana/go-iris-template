.PHONY: build
build: ## Dockerfileをキャッシュを使用せずにビルドします．
	docker-compose build --no-cache --force-rm

.PHONY: up
up: ## Dockerfileから作成したコンテナを立ち上げます．
	docker-compose up db go-iris-sample

.PHONY: down
down: ## コンテナを停止して削除します。
	docker-compose down

.PHONY: ps
ps: ## コンテナの一覧を表示します。
	docker-compose ps

generate:
	find ./_example-mvc-api/service/ -name "mock_*.go" | xargs rm -f
	go generate ./_example-mvc-api/service