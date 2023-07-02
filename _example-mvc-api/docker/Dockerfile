FROM golang:1.20-alpine

# Environment
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

# package update
RUN apk update && apk add git

# copy source
WORKDIR /go/src/go-iris-sample
COPY . .

# go.mod パッケージのインストール
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy -go=1.20
RUN go mod download

# ホットリロードの感知
CMD ["air", "-c", ".air.toml"]
