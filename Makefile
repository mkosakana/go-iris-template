.PHONY: install
install: ## go.mod 必要なパッケージのインストール
	go get github.com/kataras/iris/v12@master &&\
	go get github.com/go-sql-driver/mysql &&\
    go get github.com/go-gorp/gorp &&\
    go get github.com/kataras/iris/v12/httptest@v12.2.0-beta1.0.20220423202625-3b95c85d3db6 &&\
    go get github.com/kataras/iris/v12/websocket@v12.2.0-beta1.0.20220423202625-3b95c85d3db6 &&\
    go get github.com/kataras/iris/v12/versioning@v12.2.0-beta1.0.20220423202625-3b95c85d3db6 &&\
    go get github.com/joho/godotenv &&\
    go mod download

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

.PHONY: logs
logs: ## コンテナのログを逐次表示します。
	docker-compose logs -f
