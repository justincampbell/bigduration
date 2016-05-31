default: test lint

test: dependencies
	go test ./...

lint: dependencies
	@which gometalinter || (go get -u github.com/alecthomas/gometalinter && gometalinter --install --update)
	gometalinter --deadline 1m ./...

dependencies:
	go get -t ./...

.PHONY: default test lint dependencies
