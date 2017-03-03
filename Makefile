IMAGE_NAME=fibgo

test:
	@go test

test-all:
	@go test -cover -bench .

build-docker:
	@docker build -t $(IMAGE_NAME) .

prepare-install:
	@go get -v

install:
	@go install ./...

prepare-check:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install

check:
	@gometalinter --deadline=15s

check-all: prepare-check check
