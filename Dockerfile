FROM golang:1.8.0-alpine

WORKDIR /go/src/github.com/uudashr/fibgo

COPY . .

RUN apk update && apk upgrade && \
    apk --no-cache --update add git && \
    go get -v ./... && go install -v ./... && \
    apk del git && rm -rf /var/cache/apk/*

WORKDIR /go

EXPOSE 8080

CMD fibgo-server -port 8080
