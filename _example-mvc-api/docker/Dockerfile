FROM golang:1.21.2-alpine3.18

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN apk update &&\
    apk add --update git make

WORKDIR /go/src/app
COPY .. .

RUN go mod download &&\
    go install github.com/cosmtrek/air@latest

# ホットリロード
CMD ["air", "-c", ".air.toml"]
