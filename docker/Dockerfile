FROM golang:1.17-alpine

# appディレクトリの作成
RUN mkdir /go/src/go-iris-sample

# copy source
COPY .. /go/src/go-iris-sample/
COPY .. /usr/local/go/src/

# ワーキングディレクトリの設定
WORKDIR /go/src/go-iris-sample

# Environment
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# package update
RUN apk update && apk add git

# go.mod パッケージのインストール
RUN go get github.com/cosmtrek/air@v1.29.0
RUN go mod download

# ホットリロードの感知
CMD ["air", "-c", ".air.toml"]
