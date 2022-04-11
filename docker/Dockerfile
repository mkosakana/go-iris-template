FROM golang:1.16-alpine

# appディレクトリの作成
RUN mkdir /go/src/go-iris-sample

# Environment
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# package update
RUN apk update && apk add git

# copy source
COPY .. /go/src/go-iris-sample

# ワーキングディレクトリの設定
WORKDIR /go/src/go-iris-sample
