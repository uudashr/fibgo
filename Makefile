IMAGE_NAME=fibgo

install:
	@go install ./...

prepare-install:
	@go get -t -v ./...

test:
	@go test

test-all:
	@go test -cover -bench .

check:
	@gometalinter ./...

prepare-check:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install

docker-build:
	@docker build -t $(IMAGE_NAME) .

docker-run:
	@docker run --rm -it -p 8080:8080 $(IMAGE_NAME)

docker-console:
	@docker run --rm -it -p $(IMAGE_NAME) /bin/sh
