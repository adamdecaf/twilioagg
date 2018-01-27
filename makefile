.PHONY: build test vet

build:
	GOOS=darwin GOARCH=amd64 go build -o twilioagg-osx-amd64 github.com/adamdecaf/twilioagg
	GOOS=linux GOARCH=amd64 go build -o twilioagg-linux-amd64 github.com/adamdecaf/twilioagg

check:
	go vet ./...
	go fmt ./...

test: check
	go test ./...

deps:
	dep ensure

ci: test
