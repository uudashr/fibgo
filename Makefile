install:
	@go install ./...

prepare-install:
	@go get -v

test:
	@go test

test-all:
	@go test -cover -bench .

check:
	@gometalinter --deadline=15s

prepare-check:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install
