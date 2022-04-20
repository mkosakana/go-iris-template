FROM golang:1.17-alpine

# appディレクトリの作成
RUN mkdir /go/src/go-iris-sample

# copy source
COPY .. /go/src/go-iris-sample/
COPY .. /usr/local/go/src/

# ワーキングディレクトリの設定
WORKDIR /go/src/go-iris-sample/

# Environment
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# package update
RUN apk update && apk add git

# download packages
RUN go mod download
RUN cd _example-mvc-api/
RUN go mod tidy
